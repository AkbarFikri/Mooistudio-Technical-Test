package server

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/handler"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/repository"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/service"
	handler3 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/cart/handler"
	repository3 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/cart/repository"
	service3 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/cart/service"
	handler4 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/order/handler"
	repository4 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/order/repository"
	service4 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/order/service"
	handler2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/product/handler"
	repository2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/product/repository"
	service2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/product/service"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/config/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app      *gin.Engine
	handlers []Handler
}

type Handler interface {
	Endpoints(s *gin.RouterGroup)
}

func New(app *gin.Engine) *Server {
	s := &Server{app: app}
	db, err := database.New()
	if err != nil {
		panic(err)
	}

	if err := database.MigrateDB(db); err != nil {
		panic(err)
	}

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	productRepo := repository2.NewProductRepository(db)
	categoryRepo := repository2.NewCategoryRepository(db)
	productService := service2.NewProductService(productRepo)
	categoryService := service2.NewCategoryService(categoryRepo)
	productHandler := handler2.NewProductHandler(productService, categoryService)

	cartRepo := repository3.NewCartRepository(db)
	cartService := service3.NewCartService(cartRepo, productRepo)
	cartHandler := handler3.NewCartHandler(cartService)

	orderRepo := repository4.NewOrderRepository(db)
	orderService := service4.NewOrderService(orderRepo, cartRepo)
	orderHandler := handler4.NewOrderHandler(orderService)

	s.handlers = []Handler{authHandler, productHandler, cartHandler, orderHandler}

	return s
}

func (s *Server) SetupRoute() {
	s.app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ready!!!"})
	})
	s.app.Use(gin.Logger())

	for _, h := range s.handlers {
		h.Endpoints(s.app.Group("/api/v1"))
	}
}

func (s *Server) Run() {
	s.SetupRoute()
	s.app.Run()
}
