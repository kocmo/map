package yaml_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kocmo/map/reporters/yaml"
	"github.com/kocmo/map/spider"
)

var _ = Describe("reporters", func() {
	Describe("Yaml", func() {
		It("Executes yaml reporter", func() {
			expected := `assets: null
broken: null
children: null
links: null
name: test
url: ""
`
			result, _ := Execute(&spider.Result{Name: "test"})

			Expect(result).To(Equal(expected))
		})
	})
})
