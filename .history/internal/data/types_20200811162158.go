package db

// User type 
type User struct {
	ID int `json:"Id"`
	Name string `json:"Name"`
	Role int `json:"Role"` 
}