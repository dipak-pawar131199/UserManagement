package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"user-api/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}
		// Expect: Bearer <token>
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "Invalid Authorization token",
			})
			return
		}
		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			ctx.Abort()
			return
		}
		ctx.Set("user_Id", claims["user_id"])
		ctx.Next()
	}
}
