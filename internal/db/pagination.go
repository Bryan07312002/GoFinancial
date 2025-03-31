package db

import "time"

type PaginateOptions struct {
	Page     uint
	Take     uint
	SortBy   string
	SortDesc bool
}

type PaginateOptionsWithTimeWindowSearch struct {
	PaginateOptions
	TimeWindowSearch
}

type TimeWindowSearch struct {
	Start  time.Time `json:"from"` // default should be 'time.Now()'
	Finish time.Time `json:"to"`   // default should be zero
}

// PaginateResult holds the result of a paginated query
type PaginateResult[T any] struct {
	Data        []T    `json:"data"`
	Total       uint64 `json:"total"`
	CurrentPage uint   `json:"current_page"`
	PageSize    uint   `json:"page_size"`
	TotalPages  uint   `json:"total_pages"`
}
