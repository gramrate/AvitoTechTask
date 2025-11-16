package team

import (
	"AvitoTechTask/internal/adapters/controller/api/validator"
	"AvitoTechTask/internal/domain/dto"
	"context"

	"github.com/go-playground/form"
	"github.com/labstack/echo/v4"
)

type teamService interface {
	Create(ctx context.Context, req *dto.CreateTeamRequest) (*dto.CreateTeamResponse, error)
	GetByNameWithMembers(ctx context.Context, req *dto.GetTeamRequest) (*dto.GetTeamResponse, error)
}
type Handler struct {
	teamService teamService
	validator   *validator.Validator
	formDecoder *form.Decoder
}

func NewHandler(
	teamService teamService,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *Handler {
	return &Handler{
		teamService: teamService,
		validator:   validator,
		formDecoder: formDecoder,
	}
}

func (h *Handler) Setup(router *echo.Group) {
	router.POST("/team/add", h.Create)
	router.GET("/team/get", h.Get)
}
