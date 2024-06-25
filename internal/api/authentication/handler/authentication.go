package handler

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/dto"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/service"
	customErr "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/error"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(as service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: as,
	}
}

func (h *AuthHandler) Endpoints(s *gin.RouterGroup) {
	auth := s.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var req dto.AuthRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = customErr.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	payload, err := h.AuthService.Register(c, req)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully registered new user"),
		response.WithHttpCode(http.StatusCreated),
	).Send(ctx)
	return
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = customErr.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	payload, err := h.AuthService.Login(c, req)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithHttpCode(http.StatusOK),
		response.WithMessage("successfully login"),
	).Send(ctx)
	return
}
