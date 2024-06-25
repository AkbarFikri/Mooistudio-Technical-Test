package repository

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type CategoryRepository interface {
	Save(ctx context.Context, category domain.Category) error
	FindAll(ctx context.Context) ([]domain.Category, error)
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Save(ctx context.Context, category domain.Category) error {
	arg := map[string]interface{}{
		"id":   category.ID,
		"name": category.Name,
	}

	_, err := r.db.NamedExecContext(ctx, CreateProductCategory, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	arg := map[string]interface{}{}

	query, args, err := sqlx.Named(GetAllCategories, arg)
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

	var categories []domain.Category

	for rows.Next() {
		var category domain.Category
		err := rows.StructScan(&category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
