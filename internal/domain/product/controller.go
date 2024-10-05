package product

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	uc *UseCase
}

func NewController(engine *gin.Engine, uc *UseCase) {
	controller := &Controller{
		uc: uc,
	}

	productGroup := engine.Group("/v1/product")
	{
		productGroup.GET("/category/:id", controller.GetProductByCategoryId())
	}
}

func (c *Controller) GetProductByCategoryId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		categoryId, err := strconv.Atoi(param)
		if err != nil {
			response.NewResponse(http.StatusBadRequest, "something went wrong", nil).Send(ctx)
			return
		}

		products, err := c.uc.GetProductByCategoryId(categoryId)
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, "something went wrong", nil).Send(ctx)
			return
		}

		var productResponses []*GetProductByCategoryResponse

		for _, p := range *products {
			productResponse := &GetProductByCategoryResponse{
				Id:           p.Id,
				Name:         p.Name,
				Price:        p.Price,
				Stock:        p.Stock,
				CategoryId:   p.CategoryId,
				CategoryName: p.Category.Name,
			}

			productResponses = append(productResponses, productResponse)
		}

		response.NewResponse(http.StatusOK, "success get products", productResponses).Send(ctx)
	}
}
