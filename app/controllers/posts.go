package controllers

import (
	"github.com/revel/revel"
)

type Post struct {
	*revel.Controller
}

func (c Post) Index() revel.Result {
	return c.Todo()
}

func (c Post) Show() revel.Result {
	return c.Todo()
}

func (c Post) Store() revel.Result {
	return c.Todo()
}

func (c Post) Update() revel.Result {
	return c.Todo()
}

func (c Post) Destroy() revel.Result {
	return c.Todo()
}