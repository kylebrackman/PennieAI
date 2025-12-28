package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PennieAI/services"
)

func TestAiService(ctx *gin.Context) {
	aiService := services.NewAIService()

	result, err := aiService.Query(ctx.Request.Context(), "Say 'Hello from PennieAI!' if you can hear me, and let me know which gpt version I am talking to.", nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to connect to OpenAI",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     "OpenAI connected successfully! 🎉",
		"ai_response": result,
	})
}

func GetAiModelVersion(ctx *gin.Context) {

	modelVersion := services.GetModelVersion()

	ctx.JSON(http.StatusOK, gin.H{
		"model_version": modelVersion,
	})
}
