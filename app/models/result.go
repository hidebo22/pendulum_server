package models

import (
	"time"
)

type Result struct {
	ID        uint       `json:"id" gorm:"primary_key" gorm:"AUTO_INCREMENT"`
	HostID    string     `json:"host"`
	GuestID   string     `json:"guest"`
	LayoutID  uint       `json:"layout"`
	Replay    string     `json:"replay" gorm:"size:10000"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
