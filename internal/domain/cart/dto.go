package cart

type AddToCartItemsRequest struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type DeleteCartItemRequest struct {
	ItemId int
	UserId int
}

type AddItemResponse struct {
	Id        int `json:"id"`
	CartId    int `json:"cart_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type GetCartResponse struct {
	Id       int            `json:"id"`
	UserId   int            `json:"user_id"`
	CartItem []ItemResponse `json:"cart_item"`
}

type ItemResponse struct {
	Id        int             `json:"id"`
	CartId    int             `json:"cart_id"`
	ProductId int             `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Product   ProductResponse `json:"product"`
}

type ProductResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
