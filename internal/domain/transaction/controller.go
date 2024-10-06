package transaction

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/response"
	"github.com/MuhammadIbraAlfathar/online-store-app/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	uc *UseCase
}

func NewController(engine *gin.Engine, uc *UseCase) {
	controller := &Controller{
		uc: uc,
	}

	transactionGroup := engine.Group("/v1/transactions")
	transactionGroup.Use(middleware.AuthMiddleware())
	{
		transactionGroup.POST("", controller.CreateTransaction())
	}
}

func (c *Controller) CreateTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.MustGet("user_id").(int)
		transaction, err := c.uc.CreateTransaction(userId)
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "Error").Send(ctx)
			return
		}

		var itemTransaction []ItemTransaction
		for _, i := range transaction.Items {
			transactionResponse := &ItemTransaction{
				Id:            i.Id,
				TransactionId: i.TransactionId,
				ProductId:     i.ProductId,
				Quantity:      i.Quantity,
				Price:         i.Price,
				Product:       ItemProductResponse{i.Product.Id, i.Product.Name, i.Product.Price, i.Product.Stock, i.Product.CategoryId, i.Product.Category.Name},
			}

			itemTransaction = append(itemTransaction, *transactionResponse)
		}

		responses := &ResponseTransaction{
			Id:     transaction.Id,
			UserId: transaction.UserId,
			User: UserTransactionResponse{
				Id:      transaction.User.Id,
				Email:   transaction.User.Email,
				Address: transaction.User.Address,
			},
			TotalPrice: transaction.TotalPrice,
			Items:      itemTransaction,
			CreatedAt:  transaction.CreatedAt,
		}

		response.NewResponse(http.StatusOK, "success create transaction", responses).Send(ctx)
	}
}
