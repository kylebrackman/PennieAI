package handlers

import (
	"PennieAI/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SignUp(c *gin.Context) {
	app := config.GetFirebaseApp()
	authClient, err := app.Auth(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth client error"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
		return
	}
	//token := strings.TrimPrefix(authHeader, "Bearer ")
	////decodedToken, err := authClient.VerifyIDToken(c.Request.Context(), token)
	//
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//	return
	//}

	//email := decodedToken.Claims["email"]
	//uid := decodedToken.UID

}
