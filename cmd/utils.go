package cmd

var URL string = "http://localhost:9010/users/"

type User struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	Country string `json:"country"`
	Status  string `json:"status"`
}
