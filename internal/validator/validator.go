package validator

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func (v *Validator) MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func (v *Validator) IsInt(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func (v *Validator) GreaterThan(value string, min int) bool {
	n, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return n > min
}

func (v *Validator) IsEven(value string) bool {
	n, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return n%2 == 0
}
