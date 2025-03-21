package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Hash        string `gorm:"uniqueIndex;not null"`
	OriginalURL string `gorm:"not null"`
	Clicks      int    `gorm:"default:0"`
}
