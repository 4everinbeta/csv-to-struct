package main

import (
	"log"
	"path/filepath"
	"strings"
)

// Helper function for checking errors
// Logs the error as fatal if one occurs
func check(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func typefromfilename(filename string) (string, error) {
	var dtype string
	if strings.Contains(filename, "_") {
		dtype = strings.ToLower(strings.Split(filename, "_")[0])
	} else {
		dtype = strings.ToLower(strings.TrimSuffix(filename, filepath.Ext(filename)))
	}
	switch dtype {
	case "user":
		return dtype, nil
	case "address":
		return dtype, nil
	case "test":
		return dtype, nil
	default:
		return "", &UnsupportedType{dtype}
	}
}
