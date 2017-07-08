package controllers

import (
	"pendulum/app/models"

	"github.com/revel/revel"
)

type Result struct {
	*revel.Controller
}

type ResultInfo struct {
	Replay          string                  `json:"replay"`
	LayoutBuildings []models.LayoutBuilding `json:"buildings"`
}

func (c Result) Index(resultID uint) revel.Result {
	db := GormConnect()
	defer db.Close()

	var result models.Result
	db.Last(&result)

	var layout models.Layout
	db.First(&layout, result.LayoutID)

	var buildings []models.LayoutBuilding
	db.Model(&layout).Related(&buildings)

	var resultInfo ResultInfo
	resultInfo.Replay = result.Replay
	resultInfo.LayoutBuildings = buildings

	return c.RenderJSON(resultInfo)
}

func (c Result) Create(userID string, layoutID uint, replay string) revel.Result {
	db := GormConnect()
	db.AutoMigrate(&models.Result{})
	defer db.Close()

	result := models.Result{LayoutID: layoutID, Replay: replay}
	db.Create(&result)

	var chests []models.UserChest
	db.Where("user_id = ?", userID).Find(&chests)

	for _, chest := range chests {
		if chest.Type == 0 {
			chest.Type = 1
			//chest.UnlockedAt = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
			db.Save(&chest)
			return c.RenderJSON(chest)
		}
	}

	return c.RenderJSON(chests)
}
