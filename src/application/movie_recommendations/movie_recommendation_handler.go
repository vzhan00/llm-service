package movierecommendations

import (
	"encoding/json"
	"net/http"

	"github.com/vzhan00/llm-service/logger"
	moviedomain "github.com/vzhan00/llm-service/src/domain/movie_recommendations"
)

type MovieRecommendationHandler struct {
	movieRecommendationController *moviedomain.MovieRecommender
}

func NewMovieRecommendationHandler(movieRecommendationController moviedomain.MovieRecommender) *MovieRecommendationHandler {
	return &MovieRecommendationHandler{
		movieRecommendationController: &movieRecommendationController,
	}
}

func (handler *MovieRecommendationHandler) GetMovieRecommendations(writer http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Handling movie recommendations request - MovieRecommendationHandler")
	(*handler.movieRecommendationController).GetMovieRecommendations()
	json.NewEncoder(writer).Encode("not test")
}