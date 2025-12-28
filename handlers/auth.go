package handlers

import (
	"PennieAI/config"
	"PennieAI/models"
	"PennieAI/repository"
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
	var photoURL *string
	if photo, ok := decodedToken.Claims["picture"].(string); ok && photo != "" {
		photoURL = &photo
	}

	userEntry := models.User{
		FirebaseUID: uid,
		Email:       email.(string),
		PhotoURL:    photoURL,
	}

	err = repository.CreateUser(&userEntry)
	if err != nil {
		// Todo: error logic to delete the created user from FirebaseUI in the firebase console if the db fails to create here. Do NOT want user created in Firebase without it being in the postgres db here
		fmt.Println("Error creating user in database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User signed up successfully",
		"uid":     uid,
		"email":   email,
		"photo":   photoURL,
	})
}
