package movierecommendations

import (
	"embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/vzhan00/llm-service/logger"
)

//go:embed prompts/castle_recommendations.txt
var promptFile embed.FS

func BuildCastleMovieRecommendationPrompt(watchedMovies WatchedMovies) (string, error) {
	logger.Log.Info("Building prompt for castle movie recommendations")
	
	prompt, err := promptFile.ReadFile("prompts/castle_recommendations.txt")
	if err != nil {
		logger.Log.Error("Castle movie recommendation prompt failed to read - BuildCastleMovieRecommendationPrompt", err)
		return "", fmt.Errorf("castle movie recommendation prompt failed to read - BuildCastleMovieRecommendationPrompt: %w", err)
	}

	template, err := template.New("prompt").Parse(string(prompt))
	if err != nil {
		logger.Log.Error("Castle movie recommendation prompt template failed - BuildCastleMovieRecommendationPrompt", err)
		return "", fmt.Errorf("castle movie recommendation prompt template failed - BuildCastleMovieRecommendationPrompt: %w", err)
	}

	var parsedPrompt strings.Builder
	err = template.Execute(&parsedPrompt, watchedMovies)
	if err != nil {
		logger.Log.Error("Castle movie recommendation prompt template failed to execute - BuildCastleMovieRecommendationPrompt", err)
		return "", fmt.Errorf("castle movie recommendation prompt template failed to execute - BuildCastleMovieRecommendationPrompt: %w", err)
	}

	return parsedPrompt.String(), nil
}