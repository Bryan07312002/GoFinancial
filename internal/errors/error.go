package errors

import "encoding/json"

type ErrorByField struct {
	Errors map[string]string
}

func New(e map[string]string) ErrorByField {
	return ErrorByField{Errors: e}
}

func (e ErrorByField) Error() string {
	errorJson, err := json.Marshal(e.Errors)

	if err != nil {
		errorJson = []byte("internal server error")
	}

	return string(errorJson)
}
