package main

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Username    string `validate:"nonzero"`
	FirstName   string `validate:"nonzero"`
	LastName    string `validate:"nonzero"`
	Email       string `validate:"regexp=^[0-9a-zA-Z]+@[0-9a-zA-Z]+(\\.[0-9a-zA-Z]+)+$"`
	Phone       string `validate:"min=10"`
	DateOfBirth string
}

type Users []User

func (u User) name() string {
	return "user"
}

func (u *User) Copy() *User {
	copy := new(User)
	copy.Username = u.Username
	copy.FirstName = u.FirstName
	copy.LastName = u.LastName
	copy.Email = u.Email
	copy.Phone = u.Phone
	copy.DateOfBirth = u.DateOfBirth
	return copy
}

func (u *User) Unmarshal(reader *csv.Reader) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(u).Elem()
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

func (u *User) populateFrom(reader *csv.Reader) error {
	return u.Unmarshal(reader)
}

func (u *User) String() string {
	return fmt.Sprintf("Username: %s\tFirstName: %s\tLastName: %s\tEmail: %s\tPhone: %s\tDateOfBirth: %s", u.Username, u.FirstName, u.LastName, u.Email, u.Phone, u.DateOfBirth)
}
