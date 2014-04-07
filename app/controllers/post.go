package controllers

import (
	"github.com/antonyho/freecycle.in.net/app/models"
	"github.com/antonyho/freecycle.in.net/app/utils"
	"github.com/revel/revel"
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
	item.PostDate = now
	item.UpdateDate = now
	
	// TODO add validations

	dbutils := new(utils.DbUtils)
	session, db := dbutils.GetSession()
	defer session.Close()
	itemCollection := db.C("item")
	err := itemCollection.Insert(item)
	if err != nil {
		revel.ERROR.Println("Cannot insert record in to 'item'")
		revel.ERROR.Println(err)
	}

	// TODO Search for the tags in current post. Ff any tag is not created. insert it to database.

	return p.Render()
}

func (p *Post) List() revel.Result {
	dbutils := new(utils.DbUtils)
	session, db := dbutils.GetSession()
	defer session.Close()
	itemCollection := db.C("item")
	// TODO find all Action items ordered by date

	return p.Render()
}

/*
this function can be used by both for post owner and public
this should be requested by AJAX and response with JSON
public user should not be able to update post
*/
func (p *Post) Update(item models.Item) revel.Result {
	now := time.Now().Unix()
	item.UpdateDate = now

	return p.Render()
}

func (p *Post) ListFor(user models.User) revel.Result {
	return p.Render()
}

func (p *Post) ListBy(tag models.Tag) revel.Result {
	return p.Render()
}

func (p *Post) Request(req models.Request) revel.Result {
	return p.Render()
}
