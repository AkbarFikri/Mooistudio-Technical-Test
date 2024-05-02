package dto

type OrderResponse struct {
	ID     string       `json:"id"`
	Total  float64      `json:"total"`
	Status string       `json:"status"`
	Items  []OrderItems `json:"items"`
}

type OrderItems struct {
	ID           string  `json:"id"`
	ProductID    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Qty          int     `json:"quantity"`
}
