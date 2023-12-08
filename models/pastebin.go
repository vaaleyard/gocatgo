package models

import (
	"gorm.io/gorm"
)

type Pastebin struct {
	gorm.Model
	ShortID string `gorm:"primaryKey"`
	File    []byte
}

func (paste *Pastebin) New(db *gorm.DB) {
	db.Create(&paste)
}

func (paste *Pastebin) Get(db *gorm.DB) {
	db.First(&paste)
}

func (paste *Pastebin) GetShortID(db *gorm.DB, shortid string) {
	db.First(&paste, "short_id = ?", shortid)
}
