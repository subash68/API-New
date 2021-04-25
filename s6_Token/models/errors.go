// Package models ...
package models

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v8"
)

// DbModelError for creating custom errors resulting from Database calls
type DbModelError struct {
	ErrCode     string
	ErrTyp      string
	Err         error
	SuccessResp map[string]string
}

// DetailedErrModel ...
type DetailedErrModel struct {
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Target    string `json:"target,omitempty"`
	Condition string `json:"condition,omitempty"`
}

// RespErrModel ...
type RespErrModel struct {
	Code    string        `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Target  string        `json:"target,omitempty"`
	Errors  []DetailError `json:"errors,omitempty"`
}

type errorMessage struct {
	Code    string        `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Target  string        `json:"target,omitempty"`
	Detail  []DetailError `json:"Details,omitempty"`
}

// ErrorMessages ...
type ErrorMessages struct {
	Status string       `json:"Status"`
	Err    errorMessage `json:"Error,omitempty"`
}

// FieldError ...
type FieldError struct {
	Err validator.FieldError
}

// DetailError ...
type DetailError struct {
	Code      string `json:"code,omitempty"`
	Message   string `json:"Message,omitempty"`
	Target    string `json:"target,omitempty"`
	Condition string `json:"condition,omitempty"`
}

func (q FieldError) String() DetailError {
	var sb strings.Builder

	var errormessage DetailError
	errField := strings.ToLower(q.Err.Field[:1]) + q.Err.Field[1:]

	sb.WriteString("validation failed on field '" + errField + "'")
	sb.WriteString(", condition: " + q.Err.ActualTag)

	if q.Err.Param != "" {
		sb.WriteString(" { " + q.Err.Param + " }")
	}

	if q.Err.Value != nil && q.Err.Value != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.Err.Value))
	}

	errormessage.Code = "S1AUT002"
	errormessage.Message = sb.String()
	errormessage.Target = errField
	errormessage.Condition = q.Err.ActualTag

	return errormessage
}
