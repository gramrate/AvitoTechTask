package user

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsersPRs(c echo.Context) error {
	userIDStr := c.QueryParam("user_id")
	status := c.QueryParam("status")

	// Парсим user_id
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeBadRequest,
				Message: "Invalid user_id format",
			},
		})
	}

	// Создаем запрос вручную
	req := dto.GetUsersPRRequest{
		UserID: userID,
		Status: &status,
	}
	if status == "" {
		req.Status = nil
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
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, resp)
}
