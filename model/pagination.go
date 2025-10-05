package model

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"perPage"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"perPage"`
}
