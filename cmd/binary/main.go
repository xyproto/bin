package main

import (
	"fmt"
	"os"

	"github.com/xyproto/bin"
)

const versionString = "binary 1.0.1"

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, versionString)
		fmt.Fprintf(os.Stderr, "usage: %s [FILENAME]\n", os.Args[0])
		os.Exit(1)
	}
	filenames := os.Args[1:]
	for _, filename := range filenames {
		isBinary, err := bin.BinaryFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		if isBinary {
			if len(filenames) > 1 {
				fmt.Printf("%s:\t%s\n", filename, "binary")
			} else {
				fmt.Println("binary")
			}
		} else {
			if len(filenames) > 1 {
				fmt.Printf("%s:\t%s\n", filename, "text")
			} else {
				fmt.Println("text")
			}
		}
	}
}