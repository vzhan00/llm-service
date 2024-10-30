package movierecommendations

import (
	"encoding/json"
	"net/http"

	"github.com/vzhan00/llm-service/logger"
	moviedomain "github.com/vzhan00/llm-service/src/domain/movie_recommendations"
)

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
	var watchedMovies moviedomain.WatchedMovies

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&watchedMovies)
	if err != nil {
		logger.Log.Error("Invalid JSON input: ", err)
		http.Error(writer, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	prompt, err := moviedomain.BuildCastleMovieRecommendationPrompt(watchedMovies)
	if err != nil {
		logger.Log.Error("Prompt failed to build: ", err)
		http.Error(writer, "Prompt failed to build", http.StatusInternalServerError)
		return
	}

	logger.Log.Info(3)
	logger.Log.Info(prompt)

	recommendations, err := (*handler.movieRecommendationController).GetMovieRecommendations(prompt)
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
