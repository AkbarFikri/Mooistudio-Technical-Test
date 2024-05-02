package domain

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CategoryID  string    `db:"category_id"`
	Category    string    `db:"category_name"`
	Price       uint64    `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (p *Product) Create() {
	p.ID = uuid.NewString()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

type Category struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (c *Category) Create() {
	c.ID = uuid.NewString()
}
