package cart

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/response"
	"github.com/MuhammadIbraAlfathar/online-store-app/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	uc *UseCase
}

type AddToCartRequests struct {
	Products []AddToCartItemsRequest `json:"products"`
}

func NewController(engine *gin.Engine, uc *UseCase) {
	controller := &Controller{
		uc: uc,
	}

	cartGroup := engine.Group("/v1/cart")
	cartGroup.Use(middleware.AuthMiddleware())
	{
		cartGroup.POST("/items", controller.AddItem())
		cartGroup.GET("/items", controller.GetCartByUserId())
	}
}

func (c *Controller) AddItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("user_id").(int)

		var req AddToCartRequests
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "Error").Send(ctx)
			return
		}

		cart, err := c.uc.AddProductToCart(req.Products, userId)
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "Error").Send(ctx)
			return
		}

		var addItemResponses []*AddItemResponse

		for _, i := range *cart {
			addItemResponse := &AddItemResponse{
				Id:        i.Id,
				CartId:    i.CartId,
				ProductId: i.ProductId,
				Quantity:  i.Quantity,
			}

			addItemResponses = append(addItemResponses, addItemResponse)
		}

		response.NewResponse(http.StatusOK, "success add product to cart", addItemResponses).Send(ctx)
	}
}

func (c *Controller) GetCartByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("user_id").(int)

		cart, err := c.uc.GetCartUserByUserId(userId)
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "Error").Send(ctx)
			return
		}

		var getCartItemResponses []ItemResponse
		for _, i := range cart.CartItem {
			cartResponse := &ItemResponse{
				Id:        i.Id,
				CartId:    i.CartId,
				ProductId: i.ProductId,
				Quantity:  i.Quantity,
				Product:   ProductResponse{i.Product.Id, i.Product.Name, i.Product.Price, i.Product.Stock, i.Product.CategoryId, i.Product.Category.Name},
			}

			getCartItemResponses = append(getCartItemResponses, *cartResponse)

		}

		responses := &GetCartResponse{
			Id:       cart.Id,
			UserId:   cart.UserId,
			CartItem: getCartItemResponses,
		}

		response.NewResponse(http.StatusOK, "success get cart", responses).Send(ctx)
	}
}
