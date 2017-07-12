package controllers

import (
	"github.com/revel/revel"
)

type Post struct {
	*revel.Controller
}

func (c Post) Index() revel.Result {
	return c.Render()
}
