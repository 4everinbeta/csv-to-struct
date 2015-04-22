package main

import (
	"encoding/csv"
)

type Datatype interface {
	name() string
	Unmarshal(reader *csv.Reader) error
	String() string
}
