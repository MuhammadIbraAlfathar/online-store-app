package product

type GetProductByCategoryResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
