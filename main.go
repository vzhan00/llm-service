package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	logger "github.com/vzhan00/llm-service/logger"
	movieapp "github.com/vzhan00/llm-service/src/application/movie_recommendations"
	movieinfra "github.com/vzhan00/llm-service/src/infrastructure/movie_recommendations"
	infra "github.com/vzhan00/llm-service/src/infrastructure"
)

// RecoveryMiddleware is a custom middleware to recover from panics
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                logger.Log.Error("Recovered from panic: ", err)
                // Respond with a generic 500 error
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

func main() {
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		logger.Log.Fatalf("Failed to load .env file, error: %s", envLoadErr)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(RecoveryMiddleware)

	router.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Healthy"))
	})

	// Dependencies
	var largeLanguageModelAdapter = infra.NewLargeLanguageModelAdapter()

	var castleMovieRecommender = movieinfra.NewCastleMovieRecommender(largeLanguageModelAdapter)

	var movieRecommendationHandler = movieapp.NewMovieRecommendationHandler(castleMovieRecommender)
	var movieRecommendationRouter = movieapp.NewMovieRecommendationRouter(movieRecommendationHandler)

	router.Mount("/movie-recommendations", movieRecommendationRouter)
	
	logger.Log.Info("Starting server on port 32000")
	serverErr := http.ListenAndServe(":32000", router)
	
	if serverErr != nil {
		logger.Log.Fatalf("Failed to start server on port 32000, error: %s", serverErr)
	}
}
