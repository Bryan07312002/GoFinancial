package handlers

import (
	serviceError "financial/internal/errors"

	"net/http"
)

func writeError(err error, w http.ResponseWriter) {
	if error, ok := err.(*serviceError.ServiceError); ok {
		error.WriteJSON(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	w.Write([]byte("internal error"))
}
