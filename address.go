package main

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strconv"
)

type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	ZipCode  string
	Country  string
}

type Addresses []Address

func (a Address) name() string {
	return "address"
}

// func (a *Address) Unmarshal(record []string) {
// 	a.Address1 = record[0]
// 	a.Address2 = record[1]
// 	a.City = record[2]
// 	a.State = record[3]
// 	a.ZipCode = record[4]
// 	a.Country = record[5]
// }

func (a *Address) Copy() *Address {
	copy := new(Address)
	copy.Address1 = a.Address1
	copy.Address2 = a.Address2
	copy.City = a.City
	copy.State = a.State
	copy.ZipCode = a.ZipCode
	copy.Country = a.Country
	return copy
}

func (a *Address) Unmarshal(reader *csv.Reader) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(a).Elem()
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

func (a *Address) populateFrom(reader *csv.Reader) error {
	return a.Unmarshal(reader)
}

func (a *Address) String() string {
	return fmt.Sprintf("Address1: %s\tAddress2: %s\tCity: %s\tState: %s\tZipCode: %s\tCountry: %s", a.Address1, a.Address2, a.City, a.State, a.ZipCode, a.Country)
}
