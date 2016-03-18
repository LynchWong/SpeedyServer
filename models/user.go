package models

type User struct {
	Id int      `json:"id"`
	Name string  `json:"name"`
	Age int      `json:"age"`
}

type UserList struct {
	Users []User `json:"users"`
}