package controllers

import (
	"github.com/antonyho/freecycle.in.net/app/models"
	"github.com/antonyho/freecycle.in.net/app/utils"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"log"
	"strings"
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
	item.Status = "P"	// pending for email confirmation

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
	tags := strings.Split(item.Tags, ",")
	tagCollection := db.C("tag")
	for _, tag := range tags {
		/*
		tagResultCount, err := tagCollection.Find(bson.M{"name": tag}).Count()
		if err != nil {
			revel.WARN.Println("Cannot count the numer of selected tag: ", tag)
			revel.WARN.Println(err)
		} else {
			if tagResultCount == 0 {
				insertErr := tagCollection.Insert(models.Tag{tag})
				if insertErr != nil {
					revel.ERROR.Printf("Cannot insert tag[%s] in to 'tag'", tag)
					revel.ERROR.Println(err)
				}
			}
		}
		*/
		_, err := tagCollection.Upsert(bson.M{"name": tag}, models.Tag{tag})
		if err != nil {
			revel.ERROR.Printf("Cannot insert tag[%s] in to 'tag'", tag)
			revel.ERROR.Println(err)
		}
	}

	return p.Render()
}

func (p *Post) List() revel.Result {
	dbutils := new(utils.DbUtils)
	session, db := dbutils.GetSession()
	defer session.Close()
	itemCollection := db.C("item")
	var itemList []models.Item
	query := itemCollection.Find(bson.M{"status": "A"}).Limit(100).Sort("-postdate")
	err := query.All(&itemList)

	if err != nil {
		// TODO print error or redirect to error page
		log.Panicln(err)
	}

	// TODO Important!!! Remove owner email and ID

	return p.RenderJson(itemList)
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
