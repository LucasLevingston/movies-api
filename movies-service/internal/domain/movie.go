package domain

// Movie represents a movie entity in the system.
type Movie struct {
	ID         string `json:"id"`
	ExternalID int32  `json:"external_id"`
	Title      string `json:"title"`
	Year       string `json:"year"`
}
