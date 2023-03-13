package models

type Company struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	Name    string `json:"name"`
	Website string `json:"website"`
	Summary string `json:"summary"`
	Model
}
