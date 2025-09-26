package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func TestAiService(c *gin.Context) {
	// TODO: Review why not sending err back to client, as per all go methods
	// What is 'c' referring to above?
	godotenv.Load()

	// Create client
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)

	resp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Say 'Hello from PennieAI!' if you can hear me."),
		},
		Model: openai.ChatModelGPT4o,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to connect to OpenAI",
			"message": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     "OpenAI connected successfully! 🎉",
		"ai_response": resp.Choices[0].Message.Content,
	})

}
