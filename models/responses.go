package models

type GenericResponse struct {
	APIVersion string      `json:"api_version"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}
