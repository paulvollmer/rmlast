package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	flagInput := flag.String("i", "", "input file")
	flagErase := flag.Int("e", 0, "erase last number of lines")
	flag.Parse()

	// fmt.Println("INPUT:", *flagInput)

	d, err := ioutil.ReadFile(*flagInput)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(127)
		return
	}

	lines := strings.Split(string(d), "\n")

	for i := 0; i < len(lines)-*flagErase; i++ {
		fmt.Println(lines[i])
	}

}
