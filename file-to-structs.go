package main

import (
	"encoding/csv"
	"fmt"
	"gopkg.in/validator.v2"
	"io"
	"os"
	"strconv"
)

// Map that contains possible data types
var datatypes = map[string]int{
	"user":    1,
	"address": 2,
	"test":    3,
}

func Parsefile(file string, dtype string) ([]Datatype, error) {
	csvfile, err := os.Open(file)
	check(err)
	defer csvfile.Close()
	var results []Datatype
	reader := csv.NewReader(csvfile)
	reader.Comma = ','
	switch dtype {
	default:
		return results, &UnsupportedType{dtype}
	case "user":
		user := new(User)
		for {
			err := user.populateFrom(reader)
			if err == io.EOF {
				break
			}
			check(err)
			valid := validator.Validate(user)
			if valid == nil {
				results = append(results, user.Copy())
			} else {
				fmt.Println("Validation error?: ", valid)
			}
		}
		return results, nil
	case "address":
		addr := new(Address)
		for {
			err := addr.populateFrom(reader)
			if err == io.EOF {
				break
			}
			check(err)
			valid := validator.Validate(addr)
			if valid == nil {
				results = append(results, addr.Copy())
			} else {
				fmt.Println("Validation error?: ", valid)
			}
		}
		return results, nil
	case "test":
		test := new(Test)
		for {
			err := test.populateFrom(reader)
			if err == io.EOF {
				break
			}
			check(err)
			valid := validator.Validate(test)
			if valid == nil {
				results = append(results, test.Copy())
			} else {
				fmt.Println("Validation error?: ", valid)
			}
		}
		return results, nil
	}
}

// Error types
type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}
