package models

type UserCredentials struct {
	UserName string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
