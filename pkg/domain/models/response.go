package models

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
