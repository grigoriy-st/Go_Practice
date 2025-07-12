package handlers

import (
	"net/http"

	"myapp/middleware/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func loginHandler(c *gin.Context) {
	userID := 1

	token, err := generateToken(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func protectedHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Np token provided",
		})
		return
	}

	token, err := validateToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"]

	c.JSON(http.StatusOK, gin.H{
		"message": "Access granted",
		"user_id": userId,
	})
}
