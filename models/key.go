package models

import (
	"gorm.io/gorm"
)

type Key struct {
	gorm.Model
	Owner string
	Name  string `gorm:"unique"`
	// TransferredBy string `json:"transferred_by"` // TODO: not needed as of now
}
