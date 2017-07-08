package controllers

import (
	"pendulum/app/matchroom"
	"pendulum/app/models"

	"github.com/revel/revel"
)

type Match struct {
	*revel.Controller
}

type MatchEvents struct {
	Events []matchroom.Event `json:"events"`	
}

type MatchInfo struct {
	Opponent        matchroom.Event         `json:"opponent"`
	LayoutBuildings []models.LayoutBuilding `json:"buildings"`
}

func (c Match) Room(user string) revel.Result {
	var testMode = true

	if testMode {
		var matchInfo MatchInfo
		matchInfo.LayoutBuildings = getLayout(1)
		return c.RenderJSON(matchInfo)
	}

	subscription := matchroom.Subscribe()

	matchroom.Join(user, 0)

	// wait for something new.
	newMatch := <-subscription.New

	var matchInfo MatchInfo
	if newMatch.Host.UserID == user {
		matchInfo.Opponent = newMatch.Guest
	} else {
		matchInfo.Opponent = newMatch.Host
	}
	matchInfo.LayoutBuildings = getLayout(newMatch.Host.VP)
	return c.RenderJSON(matchInfo)
}

func getLayout(VP int) []models.LayoutBuilding {
	var layoutID int
	switch {
	case VP < 100:
		layoutID = 1
	case VP < 200:
		layoutID = 1
	case VP < 300:
		layoutID = 1
	case VP < 400:
		layoutID = 1
	case true:
		layoutID = 1
	}

	db := GormConnect()
	db.AutoMigrate(&models.Layout{}, &models.LayoutBuilding{})
	defer db.Close()

	var layout models.Layout
	db.First(&layout, layoutID)

	var buildings []models.LayoutBuilding
	db.Model(&layout).Related(&buildings)

	return buildings
}
