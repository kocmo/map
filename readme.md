> Generates map site

Simple site map generator, supports couple reporters, depth levels and etc

## Install
```sh
go get github.com/kocmo/map
```

## Usage
```sh
# Create map and output it to the terminal1
$ map http://example.com

# Create map and output map in yaml form
$ map http://example.com --reporter=yaml

# Pipe it
$ map http://example.com -r yaml > example.com.yaml

# Or use "out" flag to pipe (so you can see the spinner comparing with previous command :)
$ map http://example.com -r yaml --out=./example.com.yaml

# Define several domains
$ map http://example.com -r yaml --domains=example.net,examples.biz --out=./example.com.yaml
```
