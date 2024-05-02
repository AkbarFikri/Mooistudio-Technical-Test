package server

import (
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/authentication/handler"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/authentication/repository"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/authentication/service"
	handler3 "github.com/AkbarFikri/mooistudio_technical_test/internal/api/cart/handler"
	repository3 "github.com/AkbarFikri/mooistudio_technical_test/internal/api/cart/repository"
	service3 "github.com/AkbarFikri/mooistudio_technical_test/internal/api/cart/service"
	handler2 "github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/handler"
	repository2 "github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/repository"
	service2 "github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/service"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/config/database"
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

	s.handlers = []Handler{authHandler, productHandler, cartHandler}

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
