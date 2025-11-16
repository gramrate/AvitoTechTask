package pr

import (
	"AvitoTechTask/internal/adapters/controller/api/validator"
	"AvitoTechTask/internal/domain/dto"
	"context"

	"github.com/labstack/echo/v4"
)

type prService interface {
	Create(ctx context.Context, req *dto.CreatePRRequest) (*dto.CreatePRResponse, error)
	Merge(ctx context.Context, req *dto.MergePRRequest) (*dto.MergePRResponse, error)
	ReassignReviewer(ctx context.Context, req *dto.ReassignPRRequest) (*dto.ReassignPRResponse, error)
}
type Handler struct {
	prService prService
	validator *validator.Validator
}

func NewHandler(
	prService prService,
	validator *validator.Validator,

) *Handler {
	return &Handler{
		prService: prService,
		validator: validator,
	}
}

func (h *Handler) Setup(router *echo.Group) {
	router.POST("/pullRequest/create", h.Create)
	router.POST("/pullRequest/merge", h.Merge)
	router.POST("/pullRequest/reassign", h.Reassign)
}
