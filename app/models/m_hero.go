package models

type MHero struct {
	ID            uint    `json:"id" gorm:"primary_key"`
	Name          string  `json:"name"`
	Level         uint    `json:"level"`
	Hp            int     `json:"hp"`
	Weight        float32 `json:"weight"`
	PendulumSpeed float32 `json:"pendulumSpeed"`
}
