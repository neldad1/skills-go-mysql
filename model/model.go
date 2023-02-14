package model

// Skill is the skills table.
type Skill struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Years  int    `json:"years"`
	Rating int    `json:"rating"`
}

// Response for API.
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Skill
}
