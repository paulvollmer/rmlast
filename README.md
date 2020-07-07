# rmlast [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/paulvollmer/rmlast/blob/master/LICENSE)

cli to remove last x lines from a text source

## Installation

Install with brew

```sh
brew install paulvollmer/tap/rmlast
```

Install with go

```sh
go get -u github.com/paulvollmer/rmlast
```

## Usage

```sh
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

```sh
git clone git@github.com:paulvollmer/rmlast.git
cd rmlast
go build
```
