package dtos

type PaginatedResponse struct {
    Page       int   `json:"page"`
    PageSize   int   `json:"size"`
    Total      int64 `json:"total_items"`
    TotalPages int   `json:"total_pages"`
}