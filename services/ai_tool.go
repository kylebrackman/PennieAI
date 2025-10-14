package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"

	"PennieAI/config"
	"PennieAI/models"
)

type AIService struct {
	client openai.Client
}

// QueryOptions for functional options pattern
type QueryOptions struct {
	Schema    func(map[string]interface{}) error // Validation function
	Inferable interface{}                        // Object to link inference to
	Callback  func(*models.Inference)            // Block/yield equivalent
}

func NewAIService() *AIService {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("Please provide OPENAI_API_KEY as an environment variable")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	return &AIService{
		client: client,
	}
}

// Context below is part of Go's standard library for request-scoped values
func (s *AIService) Query(ctx context.Context, prompt string, opts *QueryOptions) (map[string]interface{}, error) {
	if opts == nil {
		opts = &QueryOptions{}
	}

	model := os.Getenv("OPENAI_MODEL_VERSION")

	// Make the API call
	response, err := s.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You are a helpful assistant for a veterinary healthcare company, Pennie. Please respond with valid JSON."),
			openai.UserMessage(prompt),
		},
		Model: os.Getenv("OPENAI_MODEL_VERSION"),
	})

	// Create inference record (equivalent to Inference.create!)
	inference := &models.Inference{
		Request:  prompt,
		Response: "", // Will be set after processing response
		Config: map[string]interface{}{
			"model":           model,
			"response_format": "json",
		},
		InferableType: nil,
		InferableID:   nil,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Handle linking to inferable object (polymorphic association)
	//if opts.Inferable != nil {
	//	// You'd implement polymorphic logic here based on type
	//	// For now, simplified:
	//	switch v := opts.Inferable.(type) {
	//	case *models.UnprocessedDocument:
	//		inference.InferableType = stringPtr("Document")
	//		inference.InferableID = &v.ID
	//	case *models.Patient:
	//		inference.InferableType = stringPtr("Patient")
	//		id := int64(v.ID)
	//		inference.InferableID = &id
	//	}
	//}

	if err != nil {
		// Log failed inference
		inference.Response = fmt.Sprintf("OpenAI API Error: %v", err)
		s.saveInference(inference)
		return nil, fmt.Errorf("OpenAI API Error: %w", err)
	}

	// Convert response to JSON string (to match Ruby storage format)
	// Todo: left off here
	responseJSON, _ := json.Marshal(map[string]interface{}{
		"choices": []map[string]interface{}{
			{
				"message": map[string]interface{}{
					"content": response.Choices[0].Message.Content,
				},
			},
		},
	})
	inference.Response = string(responseJSON)

	// Save inference to database
	s.saveInference(inference)

	// Call callback if provided (equivalent to yield)
	if opts.Callback != nil {
		opts.Callback(inference)
	}

	// Parse the actual JSON content (equivalent to your double JSON.parse)
	var parsedResponse map[string]interface{}
	if err := json.Unmarshal([]byte(response.Choices[0].Message.Content), &parsedResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	if parsedResponse == nil {
		return nil, nil
	}

	// Schema validation (equivalent to your dry-schema validation)
	if opts.Schema != nil {
		if err := opts.Schema(parsedResponse); err != nil {
			return nil, fmt.Errorf("invalid JSON response: %w", err)
		}
	}

	return parsedResponse, nil
}

// saveInference saves inference to database
// Todo: review error handling below
func (s *AIService) saveInference(inference *models.Inference) error {
	db := config.GetDB()

	configJSON, _ := json.Marshal(inference.Config)

	query := `
		INSERT INTO inferences (request, response, config, inferable_type, inferable_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id`

	/**
	The below db.Get populates inference.ID with the returned id
	db.Get is like db.QueryRow but it scans the result into the provided destination
	We pass the address of inference.ID so it gets populated
	The rest are the parameters for the SQL query
	*/
	return db.Get(&inference.ID, query,
		inference.Request,
		inference.Response,
		string(configJSON),
		inference.InferableType,
		inference.InferableID,
		inference.CreatedAt,
		inference.UpdatedAt,
	)
}

// Helper function for string pointers
func stringPtr(s string) *string {
	return &s
}

func GetModelVersion() string {
	return os.Getenv("OPENAI_MODEL_VERSION")
}
