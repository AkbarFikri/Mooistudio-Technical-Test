package domain

import (
	"github.com/google/uuid"
	"time"
)

type Cart struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	ProductID    string    `db:"product_id"`
	ProductName  string    `db:"product_name"`
	ProductPrice int64     `db:"product_price"`
	Qty          int       `db:"quantity"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (c *Cart) Create() {
	c.ID = uuid.NewString()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}
