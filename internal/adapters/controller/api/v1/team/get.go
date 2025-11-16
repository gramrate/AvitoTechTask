package team

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Get(c echo.Context) error {
	var req dto.GetTeamRequest

	if err := h.formDecoder.Decode(&req, c.QueryParams()); err != nil {
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

	resp, err := h.teamService.GetByNameWithMembers(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.ErrTeamNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeNotFound,
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
