package pr

import (
	"AvitoTechTask/internal/domain/dto"
	"AvitoTechTask/internal/domain/errorz"
	"AvitoTechTask/internal/domain/types"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Reassign(c echo.Context) error {
	var req dto.ReassignPRRequest

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

	resp, err := h.prService.ReassignReviewer(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.ErrPRNotFound) || errors.Is(err, errorz.ErrUserNotFound):
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{ // 404 Not Found
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeNotFound,
				Message: "resource not found",
			},
		})
	case errors.Is(err, errorz.ErrPRMerged):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{ // 409 Conflict
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodePRMerged,
				Message: "cannot reassign on merged PR",
			},
		})
	case errors.Is(err, errorz.ErrNotAssigned):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{ // 409 Conflict
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeNotAssigned,
				Message: "reviewer is not assigned to this PR",
			},
		})
	case errors.Is(err, errorz.ErrNoCandidate):
		return c.JSON(http.StatusConflict, dto.ErrorResponse{ // 409 Conflict
			Error: dto.ErrorDetails{
				Code:    types.ErrorCodeNoCandidate,
				Message: "no available reviewers found",
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

	return c.JSON(http.StatusOK, resp)
}
