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

type PostItem struct {
	models.Item
	PostDateString   string
	UpdateDateString string
}

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
	item.Status = "P" // pending for email confirmation

	// TODO add validations

	dbutils := new(utils.DbUtils)
	session, db := dbutils.GetSession()
	defer session.Close()
	itemCollection := db.C("item")
	err := itemCollection.Insert(item)
	if err != nil {
		revel.ERROR.Println("Cannot insert record into 'item'")
		revel.ERROR.Println(err)
	}

	// Search for the tags in current post. If any tag is not created. insert it to database.
	tags := strings.Split(item.Tags, ",")
	tagCollection := db.C("tag")
	for _, tag := range tags {
		_, err := tagCollection.Upsert(bson.M{"name": tag}, models.Tag{tag})
		if err != nil {
			revel.ERROR.Printf("Cannot insert tag[%s] in to 'tag'", tag)
			revel.ERROR.Println(err)
		}
	}

	remoteAddress := p.Controller.Request.RemoteAddr
	activityDetail := "Create new post: " + item.Title
	userActivity := models.Activity{
		PerformDate: now,
		Detail:      activityDetail,
		SourceAddr:  remoteAddress,
	}
	activityCollection := db.C("activity")
	err = activityCollection.Insert(userActivity)
	if err != nil {
		revel.ERROR.Println("Cannot insert user activity into 'activity'")
		revel.ERROR.Println(err)
	}

	return p.Render()
}

func (p *Post) List() revel.Result {
	dbutils := new(utils.DbUtils)
	session, db := dbutils.GetSession()
	if db == nil {
		return nil
	}
	defer session.Close()
	itemCollection := db.C("item")
	var itemList []models.Item
	query := itemCollection.Find(bson.M{"status": "A"}).Limit(100).Sort("-postdate")
	err := query.Select(bson.M{"email": 0, "owner": 0}).All(&itemList) // Important!!! Remove owner email and ID

	if err != nil {
		// TODO print error or redirect to error page
		log.Panicln(err)
	}

	postItemList := make([]PostItem, len(itemList))
	for idx, itm := range itemList {
		postItemList[idx] = PostItem{
			itm,
			time.Unix(itm.PostDate, 0).Format("Jan 2, 2006 at 3:04pm"),
			time.Unix(itm.UpdateDate, 0).Format("Jan 2, 2006 at 3:04pm"),
		}
	}

	// bs_grid style JSON
	var bsGridJSON interface{}
	bsGridJSON = map[string]interface{}{
		"total_rows": len(itemList),
		"page_data":  postItemList,
	}

	return p.RenderJson(bsGridJSON)
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
