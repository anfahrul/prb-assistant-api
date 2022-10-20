package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/anfahrul/prb-assistant-api/utils/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		c.Next()
	}
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(*tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("API_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err,
			})
			c.Abort()
			return
		}

		if !token.Valid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid!"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), "username", claims["username"])
		ctx = context.WithValue(ctx, "role", claims["role"])
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func DoctorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.Request.Context().Value("role")
		if role != "doctor" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "You are not doctor!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func StaffMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.Request.Context().Value("role")
		if role != "staff" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "You are not staff!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
