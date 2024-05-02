package dto

type CartRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Qty       int    `json:"qty" binding:"required"`
}

type CartResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	ProductID    string `json:"product_id"`
	Qty          int    `json:"qty"`
	ProductPrice int64  `json:"product_price"`
	ProductName  string `json:"product_name"`
}

type CartListResponse struct {
	Count int            `json:"count"`
	Total int64          `json:"total"`
	Items []CartResponse `json:"items"`
}
