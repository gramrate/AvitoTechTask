package dto

type detailsError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ErrorResponse struct {
	Error detailsError `json:"error"`
}
