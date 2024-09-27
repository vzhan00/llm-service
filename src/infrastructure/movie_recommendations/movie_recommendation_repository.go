package movierecommendations

import (
	"github.com/vzhan00/llm-service/logger"
	domain "github.com/vzhan00/llm-service/src/domain/movie_recommendations"
	infra "github.com/vzhan00/llm-service/src/infrastructure"
)

// Inherits MovieRecommender
type CastleMovieRecommender struct {
	largeLanguageModelAdapter *infra.LargeLanguageModelAdapter
}

func NewCastleMovieRecommender(largeLanguageModelAdapter *infra.LargeLanguageModelAdapter) *CastleMovieRecommender {
	return &CastleMovieRecommender{
		largeLanguageModelAdapter: largeLanguageModelAdapter,
	}
}

func (recommender *CastleMovieRecommender) GetMovieRecommendations() *[]domain.MovieRecommendation {
	logger.Log.Info("Getting movie recommendations - CastleMovieRecommender")
	prompt := "tell colin hes stupid"
	response := recommender.largeLanguageModelAdapter.GenerateContent(&prompt)

	logger.Log.Info(*response.Candidates[0].Content)
	return &[]domain.MovieRecommendation{}
}
