package controllers

import (
	"fmt"
	"log"
	"pendulum/app/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
)

type Layout struct {
	*revel.Controller
}

func (c Layout) Index(layoutID int) revel.Result {
	db := GormConnect()
	db.AutoMigrate(&models.Layout{}, &models.LayoutBuilding{})
	defer db.Close()

	var layout models.Layout
	db.Where("ID = ?", layoutID).First(&layout)

	var buildings []models.LayoutBuilding
	db.Model(&layout).Related(&buildings)

	layout.LayoutBuildings = buildings

	return c.RenderJSON(layout)
}

func (c Layout) Create() revel.Result {
	db := GormConnect()
	defer db.Close()

	var layout models.Layout
	db.Create(&layout)

	db.Last(&layout)

	return c.RenderText(fmt.Sprint(layout.ID))
}

func (c Layout) Insert(layoutID uint, buildingID uint, GridX int, GridY int) revel.Result {
	layoutBuilding := models.LayoutBuilding{LayoutID: layoutID, BuildingID: buildingID, GridX: GridX, GridY: GridY}

	db := GormConnect()
	defer db.Close()
	db.Create(&layoutBuilding)

	return c.RenderJSON(layoutBuilding)
}

func (c Layout) Update(layoutID uint, surrogateKey uint, buildingID uint, GridX int, GridY int) revel.Result {
	db := GormConnect()
	defer db.Close()

	var building models.LayoutBuilding
	db.First(&building, surrogateKey)
	if building.LayoutID == layoutID {
		db.Model(&building).Update(models.LayoutBuilding{BuildingID: buildingID, GridX: GridX, GridY: GridY})
	}
	return c.RenderJSON(building)
}

func (c Layout) Delete(layoutID uint, surrogateKey uint) revel.Result {
	db := GormConnect()
	defer db.Close()

	var building models.LayoutBuilding
	db.First(&building, surrogateKey)
	if building.LayoutID == layoutID {
		db.Delete(&building)
	}
	return c.RenderJSON(building)
}

type JwtUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	jwt.StandardClaims
}

func createTokenString() string {
	// User情報をtokenに込める
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &JwtUser{
		Name: "otiai10",
		Age:  30,
	})
	// Secretで文字列にする. このSecretはサーバだけが知っている
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
