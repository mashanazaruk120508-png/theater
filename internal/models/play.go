package models

type Play struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Genre           string `json:"genre"`
	DurationMinutes int    `json:"duration_minutes"`
	Description     string `json:"description"`
	AgeRestriction  string `json:"age_restriction"`
}
