package controllers

import (
	"pendulum/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type InitData struct {
	User      models.User        `json:"user"`
	Heros     []models.MHero     `json:"heros"`
	Pendulums []models.MPendulum `json:"pendulums"`
	Buildings []models.MBuilding `json:"buildings"`
}

func (c App) Index() revel.Result {
	return c.RenderText("test success")
}

func (c App) RegisterUser(userID string) revel.Result {
	db := GormConnect()
	db.AutoMigrate(&models.User{}, models.UserChest{})
	defer db.Close()

	user := models.User{ID: userID, Name: "name", VP: 0} //Date(1, 1, 1, 0, 0, 0, 0, time.UTC)}
	db.Create(&user)

	//宝箱スロットを追加する
	for i := 0; i < 4; i++ {
		userChest := models.UserChest{UserID: userID, Slot: i + 1, Type: 0}
		db.Create(&userChest)
	}

	return c.RenderJSON(user)
}

func (c App) GetInitData(userID string) revel.Result {
	db := GormConnect()
	db.AutoMigrate(&models.User{}, &models.UserHero{}, &models.UserPendulum{}, &models.UserChest{})
	defer db.Close()

	data := InitData{}

	var user models.User
	db.Where("ID = ?", userID).First(&user)
	db.Model(&user).Related(&user.UserHeros)
	db.Model(&user).Related(&user.UserPendulums)
	db.Model(&user).Related(&user.UserChests)
	data.User = user

	var heros []models.MHero
	db.Find(&heros)
	var pendulums []models.MPendulum
	db.Find(&pendulums)
	var buildings []models.MBuilding
	db.Find(&buildings)

	data.Heros = heros
	data.Pendulums = pendulums
	data.Buildings = buildings

	return c.RenderJSON(data)
}
