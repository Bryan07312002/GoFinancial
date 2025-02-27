package handlers

import (
	"financial/internal/db"
	"financial/internal/services"

	"errors"
	"net/http"
)

type HttpError struct {
	Status  int
	Payload string // json
}

func transalateError(err error) HttpError {
	switch {
	case errors.Is(err, services.EmailOrPasswordNotMatchError):
		return HttpError{
			Status:  401,
			Payload: err.Error(),
		}
	case err.Error() == db.ErrDuplicateEmail.Error():
		return HttpError{
			Status:  422,
			Payload: err.Error(),
		}
	}

	println(err.Error())
	return HttpError{
		Status:  500,
		Payload: "internal error",
	}
}

func ReturnError(err error, w http.ResponseWriter) {
	httpError := transalateError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpError.Status)
	w.Write([]byte(httpError.Payload))
}
