package handler

import (
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/dto"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/service"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/middleware"
	customErr "github.com/AkbarFikri/mooistudio_technical_test/internal/pkg/error"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(ps service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: ps,
	}
}

func (h *ProductHandler) Endpoints(s *gin.RouterGroup) {
	product := s.Group("product")
	product.GET("/", middleware.JwtUser(), h.GetAll)
	product.POST("/", middleware.JwtUser(), h.Create)
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	payload, err := h.ProductService.FetchProduct(c)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully find all products"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
	return
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var req dto.ProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = customErr.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	payload, err := h.ProductService.CreateProduct(c, req)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully create new product"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
	return
}
