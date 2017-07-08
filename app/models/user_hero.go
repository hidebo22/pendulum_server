package models

type UserHero struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	UserID    string `json:"userId"`
	HeroID    uint   `json:"heroId"`
	Select    int    `json:"select"`
	CardCount int    `json:"cardCount"`
}
