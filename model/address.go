package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street    string `json:"street"`
	Postal    string `json:"postal"`
	City      string `json:"city"`
	Number    string `json:"number"`
	OwnerID   uint   `json:"owner_id"`
	OwnerType string `json:"owner_type"`
}
