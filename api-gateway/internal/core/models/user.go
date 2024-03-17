package models

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
