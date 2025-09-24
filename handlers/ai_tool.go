package handlers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func TestAiService(c *gin.Context) {
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
		log.Fatal("Error:", err)
	}

	fmt.Println("✅ OpenAI connected successfully!")
	fmt.Println("Response:", resp.Choices[0].Message.Content)

}
