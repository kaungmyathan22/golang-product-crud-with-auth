package common

type ErrorResponse struct {
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
}
type PaginationParams struct {
	Page     int `query:"page,default=1"`
	PageSize int `query:"pageSize,default=10"`
}
