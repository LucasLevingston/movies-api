package domain

// Movie represents a movie in the system.
type Movie struct {
	ID         string `json:"id" example:"507f1f77bcf86cd799439011"`
	ExternalID int32  `json:"external_id" example:"8"`
	Title      string `json:"title" example:"Edison Kinetoscopic Record of a Sneeze (1894)"`
	Year       string `json:"year" example:"1894"`
}
