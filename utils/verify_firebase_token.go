package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"

	"PennieAI/config"
)

func VerifyFirebaseToken(c *gin.Context) (*auth.Token, error) {
	app := config.GetFirebaseApp()
	authClient, err := app.Auth(c.Request.Context())
	if err != nil {
		fmt.Println("Error getting auth client:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth client error"})
		return nil, err
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		fmt.Println("Missing authorization header")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
		return nil, errors.New("no authorization header")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	decodedToken, err := authClient.VerifyIDToken(c.Request.Context(), token)

	if err != nil {
		fmt.Println("Error verifying ID token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return nil, err
	}

	return decodedToken, nil
}
