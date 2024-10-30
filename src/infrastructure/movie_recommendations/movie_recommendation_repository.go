package movierecommendations

import (
	"fmt"

	"github.com/vzhan00/llm-service/logger"
	moviedomain "github.com/vzhan00/llm-service/src/domain/movie_recommendations"
	movieinfra "github.com/vzhan00/llm-service/src/infrastructure"
)

// Inherits MovieRecommender
type CastleMovieRecommender struct {
	largeLanguageModelAdapter *movieinfra.LargeLanguageModelAdapter
}

func NewCastleMovieRecommender(largeLanguageModelAdapter *movieinfra.LargeLanguageModelAdapter) *CastleMovieRecommender {
	return &CastleMovieRecommender{
		largeLanguageModelAdapter: largeLanguageModelAdapter,
	}
}

func (recommender *CastleMovieRecommender) GetMovieRecommendations(prompt string) (*[]moviedomain.MovieRecommendation, error) {
	logger.Log.Info("Getting movie recommendations - CastleMovieRecommender")
	response, err := recommender.largeLanguageModelAdapter.GenerateContent(&prompt)
	if err != nil {
		logger.Log.Error("Castle movie recommender failed to generate LLM recommendations: ", err)
		return nil, fmt.Errorf("castle movie recommender failed to generate LLM recommendations: %w", err)
	}

	logger.Log.Info(*response.Candidates[0].Content)
	return &[]moviedomain.MovieRecommendation{}, nil
}
