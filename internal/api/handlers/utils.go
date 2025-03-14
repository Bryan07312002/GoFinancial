package handlers

import (
	"financial/internal/db"
	"financial/internal/utils"
	"time"

	"net/http"
	"strconv"
)

func extractPaginateOptionsWithTimeWindowSearch(r *http.Request) db.PaginateOptionsWithTimeWindowSearch {
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
