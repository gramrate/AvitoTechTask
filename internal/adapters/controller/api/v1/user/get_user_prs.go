package user

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsersPRs(c echo.Context) error {
	var req dto.GetUsersPRRequest

	if err := h.formDecoder.Decode(&req, c.QueryParams()); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeBadRequest, // Исправлено
				Message: "Invalid request parameters",
			},
		})
	}

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeValidation, // Исправлено
				Message: "Validation failed",
			},
		})
	}

	resp, err := h.userService.GetUserPullRequests(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.ErrUserNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeNotFound, // ✅ Правильно
				Message: err.Error(),
			},
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeInternalError, // Исправлено
				Message: "Internal server error",
			},
		})
	}

	return c.JSON(http.StatusOK, resp)
}
