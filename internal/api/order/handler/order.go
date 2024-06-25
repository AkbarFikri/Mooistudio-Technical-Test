package handler

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/order/service"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/middleware"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/helper"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type OrderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(os service.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: os,
	}
}

func (h *OrderHandler) Endpoints(s *gin.RouterGroup) {
	order := s.Group("order")
	order.POST("/checkout", middleware.JwtUser(), h.Checkout)
	order.GET("/", middleware.JwtUser(), h.GetData)
	order.GET("/:id", middleware.JwtUser(), h.GetData)
}

func (h *OrderHandler) Checkout(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	user := helper.GetUserLoginData(ctx)

	payload, err := h.OrderService.CreateOrder(c, user)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully create user order"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
	return
}

func (h *OrderHandler) GetData(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	id := ctx.Param("id")
	user := helper.GetUserLoginData(ctx)

	if id == "" {
		payload, err := h.OrderService.FetchOrder(c, user)
		if err != nil {
			response.New(response.WithError(err)).Send(ctx)
			return
		}

		response.New(
			response.WithPayload(payload),
			response.WithMessage("successfully find all order record"),
			response.WithHttpCode(http.StatusOK),
		).Send(ctx)
		return

	} else {
		payload, err := h.OrderService.FetchOrderDetails(c, id)
		if err != nil {
			response.New(response.WithError(err)).Send(ctx)
			return
		}

		response.New(
			response.WithPayload(payload),
			response.WithMessage("successfully find order details record"),
			response.WithHttpCode(http.StatusOK),
		).Send(ctx)
		return
	}

}
