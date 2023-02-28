package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birthdate time.Time `json:"birthdate"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Model
}
