package models

import "time"

type UserChest struct {
	ID         uint       `json:"-" gorm:"primary_key" gorm:"AUTO_INCREMENT"`
	UserID     string     `json:"-"`
	Slot       int        `json:"slot"`
	Type       int        `json:"type"`
	UnlockedAt *time.Time `json:"unlock"`
}
