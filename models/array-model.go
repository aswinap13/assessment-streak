package models

// Array Model to parse the request
type ArrayModel struct {
	Numbers []int `json:"numbers" binding:"required"`
	Target  int   `json:"target" binding:"required,number"`
}
