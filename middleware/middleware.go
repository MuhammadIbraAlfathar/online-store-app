package middleware

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/response"
	jwttoken "github.com/MuhammadIbraAlfathar/online-store-app/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.NewResponse(http.StatusUnauthorized, "Unauthorized", nil).Send(ctx)
			ctx.Abort()
			return
		}

		tokenSlice := strings.Split(authHeader, " ")
		if len(tokenSlice) != 2 {
			response.NewResponse(http.StatusUnauthorized, "Unauthorized", nil).Send(ctx)
			ctx.Abort()
			return
		}

		tokenString := tokenSlice[1]

		token, err := jwttoken.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			response.NewResponse(http.StatusUnauthorized, "Unauthorized", nil).Send(ctx)
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.NewResponse(http.StatusUnauthorized, "Unauthorized", nil).Send(ctx)
			ctx.Abort()
			return
		}

		userId := int(claims["id"].(float64))

		ctx.Set("user_id", userId)

		ctx.Next()

	}
}
