package controllers

import (
	"github.com/revel/revel"
	"github.com/antonyho/freecycle.in.net/app/utils"
	"github.com/antonyho/freecycle.in.net/app/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Objectives() revel.Result {
	return c.Render()
}

func (c App) Tags(q string) revel.Result {
	dbutils := new(utils.DbUtils)
	session, db := dbutils.GetSession()
	defer session.Close()
	tagCollection := db.C("tag")
	var tagList []models.Tag
	var query *mgo.Query
	if len(q) > 0 {
		query = tagCollection.Find(bson.M{"name": bson.RegEx{q, "i"}}).Sort("name")
	} else {
		query = tagCollection.Find(nil).Sort("name")
	}
	
	err := query.All(&tagList)

	if err != nil {
		revel.ERROR.Println(err)
	}

	if tagList == nil {
		tagList = []models.Tag{}
	}

	return c.RenderJson(tagList)
}