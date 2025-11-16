package pr

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	var req dto.CreatePRRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeBadRequest,
				Message: "Invalid request parameters",
			},
		})
	}

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeValidation,
				Message: "Validation failed",
			},
		})
	}

	resp, err := h.prService.Create(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.ErrPRNameAlreadyUsed):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{ // 409 Conflict
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodePRExists,
				Message: "PR id already exists",
			},
		})
	case errors.Is(err, errorz.ErrUserNotFound) || errors.Is(err, errorz.ErrTeamNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{ // 404 Not Found
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeNotFound,
				Message: "resource not found",
			},
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeInternalError,
				Message: "Internal server error",
			},
		})
	}

	return c.JSON(http.StatusCreated, resp)
}
