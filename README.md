# rmlast [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/paulvollmer/rmlast/blob/master/LICENSE)
cli to remove last x lines from a text source

## Usage
```
Usage: rmlast <flags>

Flags:
  -e int
    	remove last x lines from the text source (default 0)
  -i string
    	input file
  -lb string
    	linebreak (default "\n")
  -version
    	print version and exit
```

## Development
Make shure you have installed golang.
```
git clone git@github.com:paulvollmer/rmlast.git
cd rmlast
go build
```

To release a new version install the [`gopack`](https://github.com/gobuild/gopack) tool and run:
```
gopack all
```
