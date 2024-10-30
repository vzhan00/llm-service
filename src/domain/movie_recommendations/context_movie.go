package movierecommendations

type WatchedMovies struct {
	WatchedMovies []WatchedMovie `json:"watched_movies"`
	WatchlistMovies []WatchlistMovie `json:"watchlist_movies"`
	UserPrompt string `json:"user_prompt"`
}

type WatchedMovie struct {
	Title string `json:"title"`
	Director string `json:"director"`
	Description string `json:"description"`
	Rating string `json:"rating"`
}

type WatchlistMovie struct {
	Title string `json:"title"`
	Director string `json:"director"`
	Description string `json:"description"`
}