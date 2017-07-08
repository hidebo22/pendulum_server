package models

import (
	"time"
)

type Layout struct {
	ID              uint             `json:"id" gorm:"primary_key"`
	LayoutBuildings []LayoutBuilding `json:"buildings"`
	CreatedAt       time.Time        `json:"-"`
	UpdatedAt       time.Time        `json:"-"`
	DeletedAt       *time.Time       `json:"-" sql:"index"`
}
