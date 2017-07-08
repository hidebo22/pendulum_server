package controllers

import (
	"fmt"
	"os"
	"pendulum/app/models"

	"github.com/gocarina/gocsv"
	"github.com/revel/revel"
)

type Master struct {
	*revel.Controller
}

func (c Master) Create() revel.Result {
	//CSV読み込み
	//Building
	buildingsFile, err := os.OpenFile("./m_building.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer buildingsFile.Close()

	buildings := []*models.MBuilding{}

	if err := gocsv.UnmarshalFile(buildingsFile, &buildings); err != nil { // Load buildings from file
		panic(err)
	}
	for _, building := range buildings {
		fmt.Println("building : ", building.Name)
	}
	if _, err := buildingsFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}
	//Hero
	playersFile, err := os.OpenFile("./m_player.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer playersFile.Close()

	players := []*models.MHero{}

	if err := gocsv.UnmarshalFile(playersFile, &players); err != nil { // Load players from file
		panic(err)
	}
	for _, player := range players {
		fmt.Println("player : ", player.Name)
	}
	if _, err := playersFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	//Pendulum
	pendulumsFile, err := os.OpenFile("./m_pendulum.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer pendulumsFile.Close()

	pendulums := []*models.MPendulum{}

	if err := gocsv.UnmarshalFile(pendulumsFile, &pendulums); err != nil { // Load pendulums from file
		panic(err)
	}
	for _, pendulum := range pendulums {
		fmt.Println("pendulum : ", pendulum.Name)
	}
	if _, err := pendulumsFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	db := GormConnect()
	db.Exec("drop table m_buildings;")
	db.Exec("drop table m_players;")
	db.Exec("drop table m_pendulums;")
	db.AutoMigrate(&models.MBuilding{}, &models.MHero{}, &models.MPendulum{})
	defer db.Close()

	for _, building := range buildings {
		db.Create(&building)
	}
	for _, player := range players {
		db.Create(&player)
	}
	for _, pendulum := range pendulums {
		db.Create(&pendulum)
	}

	return c.RenderText("success")
}

func (c Master) CreateTestLayout() revel.Result {
	db := GormConnect()
	db.AutoMigrate(&models.Layout{}, &models.LayoutBuilding{}, &models.MBuilding{})
	defer db.Close()

	layout := models.Layout{
		LayoutBuildings: []models.LayoutBuilding{
			{BuildingID: 001, GridX: 8, GridY: 8},
			{BuildingID: 101, GridX: 1, GridY: 2},
			{BuildingID: 102, GridX: 15, GridY: 15},
			{BuildingID: 201, GridX: 20, GridY: 15},
		},
	}
	db.Create(&layout)

	return c.Render()
}
