package dto

import "AvitoTechTask/internal/domain/types"

type ErrorDetails struct {
	Code    types.ErrorCode `json:"code"`
	Message string          `json:"message"`
}
type ErrorResponse struct {
	Error ErrorDetails `json:"error"`
}
