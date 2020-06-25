package models

type Movie struct {
	ID      string `json:"id" gorm:"type:uuid"`
	Title   string `json:"title"`
	Watched bool   `json:"watched"`
}
