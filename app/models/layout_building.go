package models

type LayoutBuilding struct {
	ID         uint      `json:"layoutItemId" gorm:"primary_key"`
	LayoutID   uint      `json:"-"`
	BuildingID uint      `json:"buildingId"`
	GridX      int       `json:"gridX"`
	GridY      int       `json:"gridY"`
	Building   MBuilding `json:"-"`
}
