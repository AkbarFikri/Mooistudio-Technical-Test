package service

import (
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/dto"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/api/product/repository"
	"github.com/AkbarFikri/mooistudio_technical_test/internal/domain"
	customErr "github.com/AkbarFikri/mooistudio_technical_test/internal/pkg/error"
	"golang.org/x/net/context"
	"log"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req dto.ProductCategoryRequest) (dto.ProductCategoryResponse, error)
	FetchCategories(ctx context.Context) ([]dto.ProductCategoryResponse, error)
}

type categoryService struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(cr repository.CategoryRepository) CategoryService {
	return &categoryService{
		CategoryRepository: cr,
	}
}

func (s categoryService) CreateCategory(ctx context.Context, req dto.ProductCategoryRequest) (dto.ProductCategoryResponse, error) {
	category := domain.Category{
		Name: req.Name,
	}
	category.Create()

	if err := s.CategoryRepository.Save(ctx, category); err != nil {
		log.Printf("error failed to save category: %v", err)
		return dto.ProductCategoryResponse{}, customErr.ErrorGeneral
	}

	return dto.ProductCategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (s categoryService) FetchCategories(ctx context.Context) ([]dto.ProductCategoryResponse, error) {
	categories, err := s.CategoryRepository.FindAll(ctx)
	if err != nil {
		log.Printf("error fetching product data from database %v", err)
		return nil, err
	}

	var res []dto.ProductCategoryResponse

	for _, c := range categories {
		res = append(res, dto.ProductCategoryResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return res, nil
}
