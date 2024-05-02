package domain

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID        string    `db:"id"`
	Total     float64   `db:"total"`
	UserID    string    `db:"user_id"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type OrderItem struct {
	ID           string `db:"id" json:"id"`
	ProductID    string `db:"product_id" json:"product_id"`
	ProductName  string `db:"product_name" json:"product_name"`
	ProductPrice int64  `db:"product_price" json:"product_price"`
	Qty          int    `db:"quantity" json:"quantity"`
	OrderID      string `db:"order_id" json:"order_id"`
}

func (o *Order) Create() {
	o.ID = uuid.NewString()
	o.Status = "waiting payment"
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
}

func (oi *OrderItem) Create() {
	oi.ID = uuid.NewString()
}
