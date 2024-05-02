package handler

import (
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/cart/dto"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/cart/service"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/middleware"
	customErr "github.com/AkbarFikri/mooistudio_technical_test/internal/pkg/error"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/pkg/helper"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type CartHandler struct {
	CartService service.CartService
}

func NewCartHandler(cs service.CartService) *CartHandler {
	return &CartHandler{
		CartService: cs,
	}
}

func (h *CartHandler) Endpoints(s *gin.RouterGroup) {
	cart := s.Group("cart")
	cart.POST("/", middleware.JwtUser(), h.Create)
	cart.GET("/", middleware.JwtUser(), h.Get)
}

func (h *CartHandler) Create(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	user := helper.GetUserLoginData(ctx)
	var req dto.CartRequest

	if err := ctx.ShouldBind(&req); err != nil {
		err = customErr.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	payload, err := h.CartService.CreateCart(c, user, req)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully create new cart for user"),
		response.WithHttpCode(http.StatusCreated),
	).Send(ctx)
	return
}

func (h *CartHandler) Get(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	user := helper.GetUserLoginData(ctx)

	payload, err := h.CartService.FetchUserCart(c, user)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}
	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully find all user cart"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
}
