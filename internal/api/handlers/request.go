package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/errors"
	"financial/internal/utils"

	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func validateRequest(data any) *errors.ServiceError {
	validate := newValidator()
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

    //fmt.Printf("%+v aaa", err)


	serviceErr := errors.BadRequestError()

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		addFieldErrors(serviceErr, validationErrors)
	} else {
		return errors.InternalServerError().WithDetails(err.Error())
	}

	return serviceErr
}

func addFieldErrors(serviceErr *errors.ServiceError, validationErrors validator.ValidationErrors) {
	for _, fieldErr := range validationErrors {
		fieldName := fieldErr.Field()
		tag := fieldErr.Tag()
		param := fieldErr.Param()

		switch tag {
		case "required":
			serviceErr.AddFieldError(
				fieldName,
				errors.RequiredField(fieldName).Message,
			)
		case "min":
			serviceErr.AddFieldError(
				fieldName,
				fmt.Sprintf("%s must be at least %s", fieldName, param),
			)
		case "max":
			serviceErr.AddFieldError(
				fieldName,
				fmt.Sprintf("%s must be at most %s", fieldName, param),
			)
		case "email":
			serviceErr.AddFieldError(
				fieldName,
				errors.InvalidFormat(fieldName).Message,
			)
		default:
			serviceErr.AddFieldError(
				fieldName,
				fmt.Sprintf("%s is invalid", fieldName),
			)
		}
	}
}

func newValidator() *validator.Validate {
	validate := validator.New()
	// register validator to use field as it is in json
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		// Handle cases where json tag might have options like ",omitempty"
		if commaIdx := strings.Index(name, ","); commaIdx != -1 {
			name = name[:commaIdx]
		}
		if name == "" {
			name = fld.Name
		}
		return name
	})

	return validate
}

func extractBody(req *http.Request, form any) error {
	if err := json.NewDecoder(req.Body).Decode(form); err != nil {
		return errors.BadRequestError().WithDetails(err.Error())
	}

	if err := validateRequest(form); err != nil {
		return err
	}

	return nil
}

func extractUserId(req *http.Request) (uint, error) {
	userID, ok := req.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		return 0, errors.UnauthorizedError()
	}

	return userID, nil
}

func extractBodyAndUserId(req *http.Request, form any) (uint, error) {
	if err := extractBody(req, form); err != nil {
		return 0, err
	}

	return extractUserId(req)
}

func extractPaginateOptionsWithTimeWindowSearch(
	r *http.Request) db.PaginateOptionsWithTimeWindowSearch {
	startAt := r.URL.Query().Get("start_at")
	finishAt := r.URL.Query().Get("finish_at")

	start := time.Now()
	if startAt != "" {
		startTime, err := utils.ParseTime(startAt)
		if err == nil {
			start = startTime
		}
	}

	finish := time.Time{}
	if finishAt != "" {
		finishTime, err := utils.ParseTime(finishAt)
		if err == nil {
			finish = finishTime
		}
	}

	paginationOptions := extractPaginationOptions(r)
	return db.PaginateOptionsWithTimeWindowSearch{
		PaginateOptions: paginationOptions,
		TimeWindowSearch: db.TimeWindowSearch{
			Start:  start,
			Finish: finish,
		},
	}
}

func extractPaginationOptions(r *http.Request) db.PaginateOptions {
	// Extract pagination options from query parameters
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	sortBy := r.URL.Query().Get("sort_by")
	sortDescStr := r.URL.Query().Get("sort_desc")

	// Set default values if not provided
	page := 1
	pageSize := 10
	sortDesc := false

	// Parse "page" query parameter
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	// Parse "page_size" query parameter
	if pageSizeStr != "" {
		parsedPageSize, err := strconv.Atoi(pageSizeStr)
		if err == nil && parsedPageSize > 0 {
			pageSize = parsedPageSize
		}
	}

	if sortDescStr != "" {
		parsedSortDescStr, err := strconv.ParseBool(sortDescStr)
		if err == nil {
			sortDesc = parsedSortDescStr
		}
	}

	return db.PaginateOptions{
		Page:     uint(page),
		Take:     uint(pageSize),
		SortBy:   sortBy,
		SortDesc: sortDesc,
	}
}
