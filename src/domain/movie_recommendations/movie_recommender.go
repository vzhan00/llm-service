package movierecommendations

type MovieRecommender interface {
	GetMovieRecommendations(prompt string) (*[]MovieRecommendation, error)
}
