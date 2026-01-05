package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"PennieAI/models"
	"PennieAI/repository"
	"PennieAI/utils"
)

func Signin(c *gin.Context) {
	// Todo: Add logic to send error if the user already exists in firebase, but possibly signed up with a different method (eg. google vs email/password)
	decodedToken, err := utils.VerifyFirebaseToken(c)
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
