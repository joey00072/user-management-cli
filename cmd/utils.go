package cmd

var URL string = "http://localhost:9010/users/"

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	Country string `json:"country"`
	Status  string `json:"status"`
}

type Custom_Error struct {
	Message string `json:"message"`
}

var user_model User
var user_models []User
var custom_error Custom_Error
var user_id int
