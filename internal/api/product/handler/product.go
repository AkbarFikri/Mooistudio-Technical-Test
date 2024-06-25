package handler

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/product/dto"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/product/service"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/middleware"
	customErr "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/error"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type ProductHandler struct {
	ProductService  service.ProductService
	CategoryService service.CategoryService
}

func NewProductHandler(ps service.ProductService, cs service.CategoryService) *ProductHandler {
	return &ProductHandler{
		ProductService:  ps,
		CategoryService: cs,
	}
}

func (h *ProductHandler) Endpoints(s *gin.RouterGroup) {
	product := s.Group("product")
	product.GET("/", middleware.JwtUser(), h.GetAll)
	product.GET("/category/:id", middleware.JwtUser(), h.GetCategory)
	product.POST("/", middleware.JwtUser(), h.Create)
	product.GET("/category", middleware.JwtUser(), h.GetCategory)
	product.POST("/category", middleware.JwtUser(), h.CreateCategory)
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

func (h *ProductHandler) GetCategory(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	payload, err := h.CategoryService.FetchCategories(c)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully find all categories"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
	return
}

func (h *ProductHandler) CreateCategory(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var req dto.ProductCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = customErr.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	payload, err := h.CategoryService.CreateCategory(c, req)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}
	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully create new product category"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
	return
}
