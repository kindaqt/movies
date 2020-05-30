package data

type Movie struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Watched bool   `json:"watched"`
}
