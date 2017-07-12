package controllers

import (
	"github.com/revel/revel"
)

type Author struct {
	*revel.Controller
}

func (c Author) Index() revel.Result {
	return c.Render()
}
