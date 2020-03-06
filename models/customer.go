package models

// Customer type details
type Customer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}
