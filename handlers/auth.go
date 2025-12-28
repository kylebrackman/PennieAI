package handlers

import (
	"PennieAI/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	app := config.GetFirebaseApp()
	authClient, err := app.Auth(c.Request.Context())

	if err != nil {
		fmt.Println("Error getting auth client:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth client error"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		fmt.Println("Missing authorization header")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	decodedToken, err := authClient.VerifyIDToken(c.Request.Context(), token)

	if err != nil {
		fmt.Println("Error verifying ID token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	email := decodedToken.Claims["email"]
	uid := decodedToken.UID
	photo := decodedToken.Claims["picture"]

	db := config.GetDB()

	_, err = db.Exec("INSERT INTO users (firebase_uid, email, photo_url) VALUES ($1, $2, $3)", uid, email, photo)

	if err != nil {
		fmt.Println("Error inserting user into database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User signed up successfully",
		"uid":     uid,
		"email":   email,
		"photo":   photo,
	})
}
