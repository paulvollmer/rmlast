# rmlast
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
