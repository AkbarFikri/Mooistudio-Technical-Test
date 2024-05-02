package dto

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	CategoryID  string `json:"category_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint64 `json:"price" binding:"required"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CategoryID  string `json:"category_id"`
	Category    string `json:"category_name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
}

type ProductCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type ProductCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
