package service

import (
	dto2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/dto"
	repository2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/cart/repository"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/order/dto"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/order/repository"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/domain"
	customErr "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/error"
	"golang.org/x/net/context"
	"log"
)

type OrderService interface {
	CreateOrder(ctx context.Context, user dto2.UserTokenData) (dto.OrderResponse, error)
	FetchOrder(ctx context.Context, user dto2.UserTokenData) ([]dto.OrderResponse, error)
	FetchOrderDetails(ctx context.Context, id string) (dto.OrderResponse, error)
}

type orderService struct {
	OrderRepository repository.OrderRepository
	CartRepository  repository2.CartRepository
}

func NewOrderService(or repository.OrderRepository, cr repository2.CartRepository) OrderService {
	return &orderService{
		OrderRepository: or,
		CartRepository:  cr,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, user dto2.UserTokenData) (dto.OrderResponse, error) {
	carts, err := s.CartRepository.FindByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("error finding cart by user id %v", err)
		return dto.OrderResponse{}, customErr.ErrorGeneral
	}

	if len(carts) == 0 {
		return dto.OrderResponse{}, customErr.ErrorNoCart
	}

	order := domain.Order{
		UserID: user.ID,
	}
	order.Create()
	var total float64

	orderItems := []domain.OrderItem{}
	orderItemsRes := []dto.OrderItems{}

	for _, c := range carts {
		dump := domain.OrderItem{
			OrderID:   order.ID,
			Qty:       c.Qty,
			ProductID: c.ProductID,
		}
		dump.Create()

		total = total + float64((int64(c.Qty) * c.ProductPrice))

		orderItems = append(orderItems, dump)

		orderItemsRes = append(orderItemsRes, dto.OrderItems{
			ID:           dump.ID,
			Qty:          dump.Qty,
			ProductID:    dump.ProductID,
			ProductName:  c.ProductName,
			ProductPrice: float64(c.ProductPrice),
		})
	}

	order.Total = total

	if err := s.OrderRepository.SaveOrder(ctx, order); err != nil {
		log.Printf("error save order: %v", err)
		return dto.OrderResponse{}, customErr.ErrorGeneral
	}

	if err := s.OrderRepository.SaveOrderItems(ctx, orderItems); err != nil {
		log.Printf("error save order items: %v", err)
		return dto.OrderResponse{}, customErr.ErrorGeneral
	}

	if err := s.CartRepository.DeleteByUserId(ctx, user.ID); err != nil {
		log.Printf("error deleting cart by user id %v", err)
		return dto.OrderResponse{}, customErr.ErrorGeneral
	}

	return dto.OrderResponse{
		ID:     order.ID,
		Total:  order.Total,
		Status: order.Status,
		Items:  orderItemsRes,
	}, nil
}

func (s *orderService) FetchOrder(ctx context.Context, user dto2.UserTokenData) ([]dto.OrderResponse, error) {
	orders, err := s.OrderRepository.FindByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("error finding order %v", err)
		return []dto.OrderResponse{}, customErr.ErrorGeneral
	}

	var res []dto.OrderResponse

	for _, o := range orders {
		res = append(res, dto.OrderResponse{
			ID:     o.ID,
			Total:  o.Total,
			Status: o.Status,
			Items:  nil,
		})
	}
	return res, nil
}

func (s *orderService) FetchOrderDetails(ctx context.Context, id string) (dto.OrderResponse, error) {
	order, err := s.OrderRepository.FindOneById(ctx, id)
	if err != nil {
		log.Printf("error finding order %v", err)
		return dto.OrderResponse{}, customErr.ErrorGeneral
	}

	orderItems, err := s.OrderRepository.FindOrderItem(ctx, order.ID)
	if err != nil {
		log.Printf("error finding order items %v", err)
		return dto.OrderResponse{}, customErr.ErrorGeneral
	}

	res := dto.OrderResponse{
		ID:     order.ID,
		Total:  order.Total,
		Status: order.Status,
	}

	items := []dto.OrderItems{}
	for _, o := range orderItems {
		items = append(items, dto.OrderItems{
			ID:           o.ID,
			Qty:          o.Qty,
			ProductID:    o.ProductID,
			ProductName:  o.ProductName,
			ProductPrice: float64(o.ProductPrice),
		})
	}

	res.Items = items

	return res, nil
}
