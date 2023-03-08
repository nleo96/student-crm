package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name" binding:"required,min=8,email"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=8"`
	Model
}
