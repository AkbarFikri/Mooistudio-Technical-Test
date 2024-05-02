package repository

import (
	"github.com/AkbarFikri/mooistudio_technical_test/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type ProductRepository interface {
	Save(ctx context.Context, product domain.Product) error
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindOne(ctx context.Context, id string) (domain.Product, error)
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Save(ctx context.Context, product domain.Product) error {
	arg := map[string]interface{}{
		"id":          product.ID,
		"name":        product.Name,
		"category_id": product.CategoryID,
		"price":       product.Price,
		"description": product.Description,
		"created_at":  product.CreatedAt,
		"updated_at":  product.UpdatedAt,
	}

	_, err := r.db.NamedExecContext(ctx, CreateProduct, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	arg := map[string]interface{}{}

	query, args, err := sqlx.Named(GetAll, arg)
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

	var products []domain.Product

	for rows.Next() {
		var product domain.Product
		err := rows.StructScan(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) FindOne(ctx context.Context, id string) (domain.Product, error) {
	arg := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(GetProductById, arg)
	if err != nil {
		return domain.Product{}, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return domain.Product{}, err
	}
	query = r.db.Rebind(query)

	var product domain.Product

	if err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&product); err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
