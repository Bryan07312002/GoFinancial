package handlers

import (
	"financial/internal/db"

	"net/http"
	"strconv"
)

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
