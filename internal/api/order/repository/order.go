package repository

import (
	"github.com/AkbarFikri/mooistudio_technical_test/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, order domain.Order) error
	SaveOrderItems(ctx context.Context, orderItems []domain.OrderItem) error
	FindByUserId(ctx context.Context, id string) ([]domain.Order, error)
	FindOneById(ctx context.Context, id string) (domain.Order, error)
	FindOrderItem(ctx context.Context, id string) ([]domain.OrderItem, error)
}

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r orderRepository) SaveOrder(ctx context.Context, order domain.Order) error {
	arg := map[string]interface{}{
		"id":         order.ID,
		"user_id":    order.UserID,
		"status":     order.Status,
		"total":      order.Total,
		"created_at": order.CreatedAt,
		"updated_at": order.UpdatedAt,
	}

	_, err := r.db.NamedExecContext(ctx, CreateOrder, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r orderRepository) SaveOrderItems(ctx context.Context, orderItems []domain.OrderItem) error {
	itemMapSlice := make([]map[string]interface{}, len(orderItems))
	for i, item := range orderItems {
		itemMapSlice[i] = map[string]interface{}{
			"id":         item.ID,
			"order_id":   item.OrderID,
			"product_id": item.ProductID,
			"quantity":   item.Qty,
		}
	}

	query, args, err := sqlx.Named(CreateOrderItem, itemMapSlice)
	if err != nil {
		return err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return err
	}
	query = r.db.Rebind(query)

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r orderRepository) FindByUserId(ctx context.Context, id string) ([]domain.Order, error) {
	arg := map[string]interface{}{
		"user_id": id,
	}

	query, args, err := sqlx.Named(GetAllByUserId, arg)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var orders []domain.Order

	for rows.Next() {
		var order domain.Order
		err := rows.StructScan(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r orderRepository) FindOneById(ctx context.Context, id string) (domain.Order, error) {
	arg := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(GetOneById, arg)
	if err != nil {
		return domain.Order{}, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return domain.Order{}, err
	}
	query = r.db.Rebind(query)

	var order domain.Order
	if err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&order); err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (r orderRepository) FindOrderItem(ctx context.Context, id string) ([]domain.OrderItem, error) {
	arg := map[string]interface{}{
		"order_id": id,
	}

	query, args, err := sqlx.Named(GetOrderItems, arg)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var orderItems []domain.OrderItem

	for rows.Next() {
		var orderItem domain.OrderItem
		err := rows.StructScan(&orderItem)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
