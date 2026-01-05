package middleware

import (
	"PennieAI/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"PennieAI/repository"
	"PennieAI/utils"
)

const UserContextKey = "authenticatedUser"

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verify the token
		decodedToken, err := utils.VerifyFirebaseToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		user, err := repository.FindUserByFirebaseUID(decodedToken.UID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		// Attach to context for handlers
		c.Set(UserContextKey, user)

		c.Next()
	}
}

// GetAuthenticatedUser retrieves the user from context
func GetAuthenticatedUser(c *gin.Context) (*models.User, bool) {
	user, exists := c.Get(UserContextKey)
	if !exists {
		return nil, false
	}
	return user.(*models.User), true
}
