package models

import (
	"time"
)

type User struct {
	ID                 string         `json:"id" gorm:"primary_key"`
	Name               string         `json:"name"`
	EcncryptedPassword []byte         `json:"-"`
	Password           string         `json:"-" sql:"-"`
	VP                 uint           `json:"vp"`
	UserHeros          []UserHero     `json:"heros"`
	UserPendulums      []UserPendulum `json:"pendulums"`
	BonusUnlockedAt    *time.Time     `json:"bonusUnlockedAt"`
	UserChests         []UserChest    `json:"chests"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          *time.Time     `json:"-"`
}
