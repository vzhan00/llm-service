package movierecommendations

import (
	"encoding/json"
	"net/http"

	"github.com/vzhan00/llm-service/logger"
	moviedomain "github.com/vzhan00/llm-service/src/domain/movie_recommendations"
)

type Prompt struct {
	Prompt string `json:"prompt"`
}

type MovieRecommendationResponse struct {
	MovieRecommendations []moviedomain.MovieRecommendation `json:"movieRecommendations"`
}

type MovieRecommendationHandler struct {
	movieRecommendationController *moviedomain.MovieRecommender
}

func NewMovieRecommendationHandler(movieRecommendationController moviedomain.MovieRecommender) *MovieRecommendationHandler {
	return &MovieRecommendationHandler{
		movieRecommendationController: &movieRecommendationController,
	}
}

func (handler *MovieRecommendationHandler) GetMovieRecommendations(writer http.ResponseWriter, request *http.Request) {
	logger.Log.Info("Handling movie recommendations request - MovieRecommendationHandler")
	var prompt Prompt

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&prompt)
	if err != nil {
		logger.Log.Error("Invalid JSON input: ", err)
		http.Error(writer, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	recommendations, err := (*handler.movieRecommendationController).GetMovieRecommendations(prompt.Prompt)
	if err != nil {
		logger.Log.Error("Failed to get movie recommendations: ", err)
		http.Error(writer, "Failed to get movie recommendations", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(MovieRecommendationResponse{
		MovieRecommendations: *recommendations,
	})
}