package infrastructure

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/vzhan00/llm-service/logger"
	"google.golang.org/api/option"
)

var largeLanguageModelApiKey = "LARGE_LANGUAGE_MODEL_API_KEY"
var modelVersion = "gemini-1.5-flash"

type ContentResponse struct {
	GeneratedResponse string
}

type LargeLanguageModelAdapter struct {
	context *context.Context
	modelClient *genai.GenerativeModel
}

func NewLargeLanguageModelAdapter() *LargeLanguageModelAdapter {
	logger.Log.Info("Initializing large language model service with model client - LargeLanguageModelAdapter")
	context := context.Background()
	client, err := genai.NewClient(context, option.WithAPIKey(os.Getenv(largeLanguageModelApiKey)))

	if err != nil {
		logger.Log.Error("Large language model service failed to retrieve client", err)
	}

	model := client.GenerativeModel(modelVersion)

	return &LargeLanguageModelAdapter{
		context: &context,
		modelClient: model,
	}
}

func (adapter *LargeLanguageModelAdapter) GenerateContent(prompt *string) *genai.GenerateContentResponse {
	logger.Log.Info("Attempting to generate content from LLM - LargeLanguageModelAdapter")
	
	response, err := adapter.modelClient.GenerateContent(*adapter.context, genai.Text(*prompt))

	if err != nil {
		logger.Log.Error("Failed to generate content from LLM: ", err)
	}

	if len(response.Candidates) == 0 {
		logger.Log.Warn("No candidate responses return from LLM")
		if response.PromptFeedback != nil {
			logger.Log.Warn("Block Reason: ", response.PromptFeedback.BlockReason)
		}
	}

	logger.Log.Info("Generated content from LLM - LargeLanguageModelAdapter")
	return response
}
