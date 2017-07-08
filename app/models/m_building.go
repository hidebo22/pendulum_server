package models

type MBuilding struct {
	ID              uint    `json:"id" gorm:"primary_key"`
	Name            string  `json:"name"`
	Level           uint    `json:"level"`
	Hp              int     `json:"hp"`
	Damage          int     `json:"damage"`
	AttackInterval  float32 `json:"attackInterval"`
	AttackRange     float32 `json:"attackRange"`
	ProjectileSpeed float32 `json:"projectileSpeed"`
}
