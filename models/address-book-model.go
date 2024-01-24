package models

import "gorm.io/gorm"

type AddressBook struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Company    string `json:"company"`
	JobTitle   string `json:"job_title"`
	Website    string `json:"website"`
	Address    string `json:"address"`
	Note       string `json:"note"`
}
