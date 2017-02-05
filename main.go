package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const version = "0.1.0"

func usage()  {
	fmt.Printf("Usage: rmlast <flags>\n\n")
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func main() {
	flagInput := flag.String("i", "", "input file")
	flagLineBreak := flag.String("lb", "\n", "linebreak")
	flagRemoveLines := flag.Int("e", 0, "remove last x lines from the text source")
	flagVersion := flag.Bool("version", false, "print version and exit")
	flag.Usage = usage
	flag.Parse()

	if *flagVersion {
			fmt.Println(version)
			os.Exit(0)
	}

	if *flagInput == "" {
		usage()
		os.Exit(127)
	}
	d, err := ioutil.ReadFile(*flagInput)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(127)
		return
	}

	lines := strings.Split(string(d), *flagLineBreak)
	linesTotal := len(lines)

	if linesTotal != 0 && *flagRemoveLines > 0 {
		for i := 0; i < linesTotal-*flagRemoveLines; i++ {
			fmt.Println(lines[i])
		}
	}
}
