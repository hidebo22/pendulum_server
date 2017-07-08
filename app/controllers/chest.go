package controllers

import (
	"github.com/revel/revel"
)

type Chest struct {
	*revel.Controller
}

func (c Chest) OpenChest(userID string, slot int) revel.Result {
	db := GormConnect()
	defer db.Close()

	return c.RenderText("")
}
