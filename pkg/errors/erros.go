package erros

import "net/http"

type APIError struct {
	Code   string `json:"code"`
	Text   string `json:"text"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

type CustomError string

func (err CustomError) Error() string {
	return string(err)

}

type BadRequestToAPI string

func (err BadRequestToAPI) Error() string {
	return string(err)
}

type ResourceNotFound string

func (err ResourceNotFound) Error() string {
	return string(err)
}

func CustomBadRequest(code string, err error) APIError {
	return APIError{
		Code:   code,
		Text:   "Bad Request",
		Status: http.StatusBadRequest,
		Detail: err.Error(),
	}
}

func BadRequest(detail string) APIError {
	return APIError{
		Code:   "bad_request",
		Text:   "Bad Request",
		Status: http.StatusBadRequest,
		Detail: detail,
	}
}

func InternalServerError(err error) APIError {
	return APIError{
		Code:   "internal_server_error",
		Text:   "Internal Server Error",
		Status: http.StatusInternalServerError,
		Detail: err.Error(),
	}
}

func Unauthorized(detail string) APIError {
	return APIError{
		Code:   "unauthorized",
		Text:   "unauthorized",
		Status: http.StatusUnauthorized,
		Detail: detail,
	}
}
