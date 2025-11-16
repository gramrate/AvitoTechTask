package team

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	var req dto.CreateTeamRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeBadRequest,
				Message: err.Error(),
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

	resp, err := h.teamService.Create(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.ErrTeamNameAlreadyUsed):
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeTeamExists,
				Message: err.Error(),
			},
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeInternalError,
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusCreated, resp)
}
