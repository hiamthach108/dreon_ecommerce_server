package constants

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

func (e AppError) Error() string {
	return e.Reason
}

func (e AppError) ToEchoHTTPError() *echo.HTTPError {
	b, _ := json.Marshal(e)
	var resp map[string]interface{}
	json.Unmarshal(b, &resp)
	return echo.NewHTTPError(e.Code, resp)
}

func NewBadRequest(err error, message string) *AppError {
	reason := ""
	if err != nil {
		reason = err.Error()
	}
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
		Reason:  reason,
	}
}

func NewInternal(err error, message string) *AppError {
	reason := ""
	if err != nil {
		reason = err.Error()
	}
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Reason:  reason,
	}
}

func NewNotFound(err error, message string) *AppError {
	reason := ""
	if err != nil {
		reason = err.Error()
	}
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
		Reason:  reason,
	}
}

func NewUnAuthorize(err error, message string) *AppError {
	reason := ""
	if err != nil {
		reason = err.Error()
	}
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
		Reason:  reason,
	}
}

func NewForbidden(err error, message string) *AppError {
	reason := ""
	if err != nil {
		reason = err.Error()
	}
	return &AppError{
		Code:    http.StatusForbidden,
		Message: message,
		Reason:  reason,
	}
}
