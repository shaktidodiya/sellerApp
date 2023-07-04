package errors

import (
	"net/http"
	"strings"
)

type APIError interface {
	APIError() (int, string)
}

type InvalidParam struct {
	Param []string
}

func (i InvalidParam) Error() string {
	return "Invalid Parameters " + strings.Join(i.Param, ",")
}

func (i InvalidParam) APIError() (int, string) {
	return http.StatusBadRequest, i.Error()
}

type DBError struct {
	Err error
}

func (d DBError) Error() string {
	return d.Err.Error()
}

func (d DBError) APIError() (int, string) {
	return http.StatusInternalServerError, d.Error()
}
