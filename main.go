package main

import (
	"flag"
	"fmt"
)

func main() {
	var source string
	flag.StringVar(&source, "file", "Test.csv", "the file to parse")
	flag.Parse()
	dtype, err := typefromfilename(source)
	check(err)
	structs, err := Parsefile(source, dtype)
	check(err)
	for i := 0; i < len(structs); i++ {
		test := structs[i]
		fmt.Printf("Results: %+v\n", test.String())
	}
}
