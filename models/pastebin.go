package models

import (
	"gorm.io/gorm"
)

type Pastebin struct {
	gorm.Model
	ShortID string `gorm:"primaryKey"`
	File    string
}

func (paste *Pastebin) New(db *gorm.DB) {
	db.Create(&paste)
}
