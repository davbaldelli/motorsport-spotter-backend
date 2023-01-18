package models

type Role string

const (
	Admin Role = "admin"
	Base  Role = "base"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Token struct {
	Role        string `json:"role"`
	Username    string `json:"username"`
	TokenString string `json:"token"`
}
