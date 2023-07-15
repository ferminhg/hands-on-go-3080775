// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("💣 recovered from panic:", err)
			log.Fatal(err)
		}
	}()
	//panic("wop wop")

	// use os package to read the file specified as a command line argument
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	if len(argsWithoutProg) != 1 {
		fmt.Println("✋ Provide a file to read")
	}

	// convert the bytes to a string
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	data := string(bs)

	// initialize a map to store the counts
	analysis := make(map[string]int)
	fmt.Println(analysis)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(data)
	//fmt.Println(words)

	// capture the length of the words slice
	analysis["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, w := range words {
		for _, c := range w {
			if unicode.IsLetter(c) {
				analysis["letters"]++
			} else if unicode.IsNumber(c) {
				analysis["numbers"]++
			} else {
				analysis["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(analysis)
}
