package models

type Base64 struct {
	Base64Str string `json:"base64" binding:"required"`
}
