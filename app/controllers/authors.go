package controllers

import (
	"github.com/revel/revel"
	"github.com/zxp86021/ChuCooBlog-golang/app/models"
)

type Author struct {
	*revel.Controller
}

func (c Author) Index() revel.Result {
	return c.Todo()
}

func (c Author) Show() revel.Result {
	return c.Todo()
}

func (c Author) Store(author models.Author) revel.Result {
    /*   
     *   username
     *   password
     *   name
     *   gender
     *   address
     */
	//var jsonData map[string]interface{}

	//c.Params.BindJSON(&author)

	//author.Password = jsonData["password"]
	c.Validation.Required(author.Username)
	c.Validation.Required(author.Password)
	c.Validation.Required(author.Name)
	c.Validation.Required(author.Gender)
	c.Validation.Required(author.Address)

	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Response.Status = 400

		return c.RenderJSON(author.Password)
	}

	c.Response.Status = 201

	return c.RenderJSON(author)
}

func (c Author) Update() revel.Result {
	return c.Todo()
}
