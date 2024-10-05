package auth

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	uc *UseCase
}

func NewController(engine *gin.Engine, uc *UseCase) {
	controller := &Controller{uc: uc}

	authGroup := engine.Group("/v1/auth")
	{
		authGroup.POST("/register", controller.Register())
	}
}

func (c *Controller) Register() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req RegisterRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "Error").Send(ctx)
			return
		}

		if err := c.uc.Register(&req); err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "Error").Send(ctx)
			return
		}

		response.NewResponse(http.StatusCreated, "Success register account", "-").Send(ctx)
	}
}
