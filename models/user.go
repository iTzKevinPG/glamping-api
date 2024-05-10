package models

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterCredentials struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Id          int    `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	PayMethodId int    `json:"payMethodId"`
}
