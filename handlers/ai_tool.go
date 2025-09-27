package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PennieAI/services"
)

func TestAiService(c *gin.Context) {
	// Create service
	aiService := services.NewAIService()

	// Call the service (equivalent to Openai.query in Ruby)
	result, err := aiService.Query(c.Request.Context(), "Say 'Hello from PennieAI!' if you can hear me, and let me know which gpt version I am talking to.", nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to connect to OpenAI",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     "OpenAI connected successfully! 🎉",
		"ai_response": result,
	})
}

func GetAiModelVersion(c *gin.Context) {

	modelVersion := services.GetModelVersion()

	c.JSON(http.StatusOK, gin.H{
		"model_version": modelVersion,
	})
}
