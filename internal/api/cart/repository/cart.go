package repository

import (
	"github.com/AkbarFikri/mooistudio_technical_test/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type CartRepository interface {
	Save(ctx context.Context, cart domain.Cart) error
	FindByUserId(ctx context.Context, id string) ([]domain.Cart, error)
}

type cartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r cartRepository) Save(ctx context.Context, cart domain.Cart) error {
	arg := map[string]interface{}{
		"id":         cart.ID,
		"user_id":    cart.UserID,
		"product_id": cart.ProductID,
		"quantity":   cart.Qty,
		"created_at": cart.CreatedAt,
		"updated_at": cart.UpdatedAt,
	}

	_, err := r.db.NamedExecContext(ctx, CreateCart, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r cartRepository) FindByUserId(ctx context.Context, id string) ([]domain.Cart, error) {
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

	var carts []domain.Cart

	for rows.Next() {
		var cart domain.Cart
		err := rows.StructScan(&cart)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
