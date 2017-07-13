package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) PostLogin() revel.Result {
	return c.Todo()
}

func (c App) PostLogout() revel.Result {
	return c.Todo()
}

func (c App) GetLogin() revel.Result {
	return c.Todo()
}