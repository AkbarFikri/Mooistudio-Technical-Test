package service

import (
	dto2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/dto"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/cart/dto"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/cart/repository"
	repository2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/product/repository"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/domain"
	customErr "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/error"
	"golang.org/x/net/context"
	"log"
)

type CartService interface {
	CreateCart(ctx context.Context, user dto2.UserTokenData, req dto.CartRequest) (dto.CartResponse, error)
	FetchUserCart(ctx context.Context, user dto2.UserTokenData) (dto.CartListResponse, error)
	DeleteCart(ctx context.Context, id string) error
}

type cartService struct {
	CartRepository    repository.CartRepository
	ProductRepository repository2.ProductRepository
}

func NewCartService(cr repository.CartRepository, pr repository2.ProductRepository) CartService {
	return &cartService{
		CartRepository:    cr,
		ProductRepository: pr,
	}
}

func (c cartService) CreateCart(ctx context.Context, user dto2.UserTokenData, req dto.CartRequest) (dto.CartResponse, error) {
	product, err := c.ProductRepository.FindOne(ctx, req.ProductID)
	if err != nil {
		log.Printf("error failed to find product %v", err)
		return dto.CartResponse{}, customErr.ErrorNotFound
	}

	cart := domain.Cart{
		ProductID: product.ID,
		UserID:    user.ID,
		Qty:       req.Qty,
	}
	cart.Create()

	if err := c.CartRepository.Save(ctx, cart); err != nil {
		log.Printf("error failed to save cart %v", err)
		return dto.CartResponse{}, customErr.ErrorGeneral
	}

	return dto.CartResponse{
		ID:        cart.ID,
		ProductID: cart.ProductID,
		UserID:    cart.UserID,
		Qty:       cart.Qty,
	}, nil
}

func (c cartService) FetchUserCart(ctx context.Context, user dto2.UserTokenData) (dto.CartListResponse, error) {
	carts, err := c.CartRepository.FindByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("error failed to find cart %v", err)
		return dto.CartListResponse{}, customErr.ErrorNotFound
	}

	var cartResponses []dto.CartResponse
	var total int64

	for _, c := range carts {
		cartResponses = append(cartResponses, dto.CartResponse{
			ID:           c.ID,
			ProductID:    c.ProductID,
			UserID:       c.UserID,
			Qty:          c.Qty,
			ProductName:  c.ProductName,
			ProductPrice: c.ProductPrice,
		})
		total = total + (int64(c.Qty) * c.ProductPrice)
	}
	return dto.CartListResponse{
		Total: total,
		Count: len(cartResponses),
		Items: cartResponses,
	}, nil
}

func (c cartService) DeleteCart(ctx context.Context, id string) error {
	cart, err := c.CartRepository.FindOneById(ctx, id)
	if err != nil {
		log.Printf("error failed to find cart %v", err)
		return customErr.ErrorNotFound
	}

	if err := c.CartRepository.DeleteById(ctx, cart.ID); err != nil {
		log.Printf("error failed to delete cart %v", err)
		return customErr.ErrorGeneral
	}

	return nil
}
