package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"

	"PennieAI/config"
	"PennieAI/models"
	"PennieAI/repository"
)

func Signin(c *gin.Context) {

	decodedToken, err := verifyFirebaseToken(c)
	if err != nil {
		c.JSON(400, gin.H{})
	}

	firebaseUID := decodedToken.UID
	email := decodedToken.Claims["email"]
	var photoURL *string
	if photo, ok := decodedToken.Claims["picture"].(string); ok && photo != "" {
		photoURL = &photo
	}

	userExists, err := repository.FindUserByFirebaseUID(firebaseUID)
	if err != nil {
		fmt.Println("Error checking user existence:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if !userExists {

		userEntry := models.User{
			FirebaseUID: firebaseUID,
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
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User signed up successfully",
		"uid":     firebaseUID,
		"email":   email,
		"photo":   photoURL,
	})
}

func verifyFirebaseToken(c *gin.Context) (*auth.Token, error) {
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
