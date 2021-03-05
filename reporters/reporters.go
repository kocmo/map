// Package reporters provides various reporter executioners
package reporters

import (
	"github.com/go-errors/errors"

	"github.com/kocmo/map/spider"

	// Reporters
	"github.com/kocmo/map/reporters/json"
	"github.com/kocmo/map/reporters/yaml"
)

// Available list of reporters
var Available = map[string]bool{
	"json": true,
	"yaml": true,
}

// Exist finds out if such reporter exist
func Exist(name string) (ok bool) {
	_, ok = Available[name]

	return
}

// Execute gets the requested reporter and feeds data to it
func Execute(name string, data *spider.Result) (string, error) {
	if data == nil {
		return "", nil
	}

	if name == "json" {
		return json.Execute(data)
	}

	if name == "yaml" {
		return yaml.Execute(data)
	}

	return "", errors.New(name + " reporter doesn't exist")
}
