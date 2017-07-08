package models

type MPendulum struct {
	ID              uint    `json:"id" gorm:"primary_key"`
	Name            string  `json:"name"`
	Level           uint    `json:"level"`
	IsUp            bool    `json:"isUp"`
	Damage          int     `json:"damage"`
	AttackInterval  float32 `json:"attackInterval"`
	AttackRange     float32 `json:"attackRange"`
	ProjectileSpeed float32 `json:"projectileSpeed"`
}
