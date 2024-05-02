package service

import (
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/dto"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/repository"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/domain"
	customErr "github.com/AkbarFikri/mooistudio_technical_test/internal/pkg/error"
	"golang.org/x/net/context"
	"log"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req dto.ProductRequest) (dto.ProductResponse, error)
	FetchProduct(ctx context.Context) ([]dto.ProductResponse, error)
}

type productService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productService{
		ProductRepository: pr,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req dto.ProductRequest) (dto.ProductResponse, error) {
	product := domain.Product{
		Name:        req.Name,
		CategoryID:  req.CategoryID,
		Price:       req.Price,
		Description: req.Description,
	}
	product.Create()

	if err := s.ProductRepository.Save(ctx, product); err != nil {
		log.Printf("error saving product data to database %v", err)
		return dto.ProductResponse{}, customErr.ErrorGeneral
	}

	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Description: product.Description,
	}, nil
}

func (s *productService) FetchProduct(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.ProductRepository.FindAll(ctx)
	if err != nil {
		log.Printf("error fetching product data from database %v", err)
		return nil, err
	}

	var res []dto.ProductResponse

	for _, p := range products {
		res = append(res, dto.ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			CategoryID:  p.CategoryID,
			Category:    p.Category,
			Price:       p.Price,
			Description: p.Description,
		})
	}

	return res, nil
}
