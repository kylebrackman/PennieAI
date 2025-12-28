package handlers

import (
	"PennieAI/config"
	//"PennieAI/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

	db := config.GetDB()

	//user := models.User{
	//	FirebaseUID: uid,
	//	Email:       email.(string),
	//	PhotoURL:    photoURL,
	//}

	_, err = db.Exec("INSERT INTO users (firebase_uid, email, photo_url) VALUES ($1, $2, $3)", uid, email, photoURL)

	if err != nil {
		fmt.Println("Error inserting user into database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User signed up successfully",
		"uid":     uid,
		"email":   email,
		"photo":   photoURL,
	})
}
