package movierecommendations

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewMovieRecommendationRouter(movieRecommendationHandler *MovieRecommendationHandler) http.Handler {
	router := chi.NewRouter()

	router.Get("/", movieRecommendationHandler.GetMovieRecommendations)

	return router
}
