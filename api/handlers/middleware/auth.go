package middleware

import (
	"crudgin/pkg/utils/jwttoken"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		log.Println(tokenStr)

		login, valid, err := jwttoken.ValidToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
			c.Abort()
			return
		}

		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("login", login)

		c.Next()
	}
}
