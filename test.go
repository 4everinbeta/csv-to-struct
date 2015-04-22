package main

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strconv"
)

type Test struct {
	Name    string
	Surname string
	Age     int64
}

func (t Test) name() string {
	return "test"
}

func (t *Test) Copy() *Test {
	copy := new(Test)
	copy.Name = t.Name
	copy.Surname = t.Surname
	copy.Age = t.Age
	return copy
}

func (t *Test) Unmarshal(reader *csv.Reader) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(t).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int", "int64":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}

func (t *Test) populateFrom(reader *csv.Reader) error {
	return t.Unmarshal(reader)
}

func (t *Test) String() string {
	return fmt.Sprintf("Name: %s\tSurname: %s\tAge: %d", t.Name, t.Surname, t.Age)
}
