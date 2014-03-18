package controllers

import (
	"github.com/antonyho/freecycle.in.net/models"
	"github.com/revel/revel"
)

type Post struct {
	*revel.Controller
}

func (p *Post) New(models.Item) revel.Result {
	return p.Render()
}

func (p *Post) List() revel.Result {
	return p.Render()
}

func (p *Post) Update(models.Item) revel.Result {
	return p.Render()
}

/*
this function can be used by both for post owner and public
this should be requested by AJAX and response with JSON
public user should not be able to update post
*/
func (p *Post) ListFor(models.User) revel.Result {
	return p.Render()
}

func (p *Post) Request(models.Request) revel.Result {
	return p.Render()
}
