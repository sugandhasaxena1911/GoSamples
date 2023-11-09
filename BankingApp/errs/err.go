package errs

import "net/http"

type AppError struct {
	Code   int    `json:"code,omitempty"`
	Errmsg string `json:"message"`
}

func (ae AppError) AsMessage() *AppError {
	return &AppError{Errmsg: ae.Errmsg}
}

func NewNotFoundError(message string) *AppError {

	return &AppError{Errmsg: message, Code: http.StatusNotFound}

}
func NewUnexpectedError(message string) *AppError {

	return &AppError{Errmsg: message, Code: http.StatusInternalServerError}

}
func NewBadError(message string) *AppError {

	return &AppError{Errmsg: message, Code: http.StatusBadRequest}

}
