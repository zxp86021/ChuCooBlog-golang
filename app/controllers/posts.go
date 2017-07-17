package controllers

import (
	"github.com/revel/revel"
	"encoding/json"
	"io/ioutil"
)

type Post struct {
	*revel.Controller
}

type PostInput struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
}

func (c Post) Index() revel.Result {
	file, _ := ioutil.ReadFile(revel.BasePath + "/storage/posts.json")

	post_data := []byte(file)

	var json_post_data []interface {}

	json.Unmarshal(post_data, &json_post_data)



	return c.RenderJSON(json_post_data)
}

func (c Post) Show() revel.Result {
	return c.Todo()
}

func (c Post) Store(input PostInput) revel.Result {
	if c.Session["username"] == "" {
		var errors []interface{}

		json.Unmarshal([]byte("[{\"Message\": \"請先登入\"}]"), &errors)

		c.Response.Status = 401

		return c.RenderJSON(Errors{errors})
	}

	c.Params.BindJSON(&input)

	//file, err := ioutil.ReadFile(revel.BasePath + "/storage/posts.json")

	//check(err)

	//post_data := []byte(file)

	return c.RenderJSON(input)
}

func (c Post) Update() revel.Result {
	if c.Session["username"] == "" {
		var errors []interface{}

		json.Unmarshal([]byte("[{\"Message\": \"請先登入\"}]"), &errors)

		c.Response.Status = 401

		return c.RenderJSON(Errors{errors})
	}

	return c.Todo()
}

func (c Post) Destroy() revel.Result {
	if c.Session["username"] == "" {
		var errors []interface{}

		json.Unmarshal([]byte("[{\"Message\": \"請先登入\"}]"), &errors)

		c.Response.Status = 401

		return c.RenderJSON(Errors{errors})
	}

	return c.Todo()
}