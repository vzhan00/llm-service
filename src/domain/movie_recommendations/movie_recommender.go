package movierecommendations

type MovieRecommender interface {
	GetMovieRecommendations() *[]MovieRecommendation
}
