package handler

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/cart/dto"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/cart/service"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/middleware"
	customErr "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/error"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/helper"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/response"
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
	cart.DELETE("/:id", middleware.JwtUser(), h.Delete)
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

func (h *CartHandler) Delete(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	id := ctx.Param("id")
	if id == "" {
		err := customErr.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	if err := h.CartService.DeleteCart(c, id); err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithHttpCode(http.StatusOK),
		response.WithMessage("successfully delete cart"),
	).Send(ctx)
}
