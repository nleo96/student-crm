package models

type Contact struct {
	ID        uint    `json:"id" gorm:"primarykey"`
	Name      string  `json:"name" bindin:"required"`
	Position  string  `json:"position"`
	Email     string  `json:"email" binding:"email"`
	Phone     string  `json:"phone"`
	CompanyID int     `json:"-"`
	Company   Company `json:"company"`
	Model
}
