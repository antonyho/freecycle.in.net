package controllers

import (
	"github.com/antonyho/freecycle.in.net/app/models"
	"github.com/antonyho/freecycle.in.net/app/utils"
	"github.com/revel/revel"
	"log"
	"time"
)

type Post struct {
	*revel.Controller
}

func (p *Post) NewForm() revel.Result {
	return p.Render()
}

func (p *Post) New(item models.Item) revel.Result {
	now := time.Now().Unix()
	log.Println(p.Params)
	item.PostDate = now
	item.UpdateDate = now
	dbutils := new(utils.DbUtils)
	session, _ := dbutils.GetSession()
	defer session.Close()
	itemCollection := session.DB("freecycle").C("item")
	err := itemCollection.Insert(item)
	if err != nil {
		revel.ERROR.Println("Cannot insert record in to 'item'")
		revel.ERROR.Println(err)
	}

	return p.Render()
}

func (p *Post) List() revel.Result {
	return p.Render()
}

/*
this function can be used by both for post owner and public
this should be requested by AJAX and response with JSON
public user should not be able to update post
*/
func (p *Post) Update(models.Item) revel.Result {
	return p.Render()
}

func (p *Post) ListFor(models.User) revel.Result {
	return p.Render()
}

func (p *Post) Request(models.Request) revel.Result {
	return p.Render()
}
