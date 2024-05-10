package models

type GenericResponse struct {
	APIVersion string      `json:"api_version"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}

type UserResponse struct {
	Id          int    `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	PayMethodId *int   `json:"payMethodId"`
}
