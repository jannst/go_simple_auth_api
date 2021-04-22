package util

import "github.com/go-openapi/strfmt"

func CreateMail(value string) *strfmt.Email{
	tmp := strfmt.Email(value)
	return &tmp
}
