package user

import (
	"AvitoTechTask/internal/adapters/controller/api/validator"
	"AvitoTechTask/internal/domain/dto"
	"context"

	"github.com/go-playground/form"
	"github.com/labstack/echo/v4"
)

type userService interface {
	GetUserPullRequests(ctx context.Context, req *dto.GetUsersPRRequest) (*dto.GetUsersPRResponse, error)
	UpdateActivity(ctx context.Context, req *dto.SetUserActivityRequest) (*dto.SetUserActivityResponse, error)
}
type Handler struct {
	userService userService
	validator   *validator.Validator
	formDecoder *form.Decoder
}

func NewHandler(
	userService userService,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *Handler {
	return &Handler{
		userService: userService,
		validator:   validator,
		formDecoder: formDecoder,
	}
}

func (h *Handler) Setup(router *echo.Group) {
	router.POST("/users/setIsActive", h.SetActive)
	router.GET("/users/getReview", h.GetUsersPRs)
}
