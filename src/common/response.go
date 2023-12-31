package common

type ErrorResponse struct {
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
}
type PaginationParams struct {
	Page     int64 `query:"page,default=1"`
	PageSize int64 `query:"pageSize,default=10"`
}
