package common

type ErrorResponse struct {
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
}
