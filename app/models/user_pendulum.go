package models

type UserPendulum struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	UserID     string `json:"userId"`
	PendulumID uint   `json:"pendulumId"`
	Select     int    `json:"select"`
	CardCount  int    `json:"cardCount"`
}
