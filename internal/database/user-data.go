package database

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistration struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
