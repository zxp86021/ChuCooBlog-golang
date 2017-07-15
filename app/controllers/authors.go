package controllers

import (
	"github.com/revel/revel"
	"github.com/zxp86021/ChuCooBlog-golang/app/models"
	"regexp"
	"io/ioutil"
	"encoding/json"
)

type Author struct {
	*revel.Controller
}

type Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (c Author) Index() revel.Result {
	if c.Session["username"] == "" {
		var errors []interface {}

		json.Unmarshal([]byte("[{\"Message\": \"請先登入\"}]"), &errors)

		c.Response.Status = 404

		return c.RenderJSON(Errors{errors})
	}

	file, _ := ioutil.ReadFile(revel.BasePath + "/storage/authors.json")

	author_data := []byte(file)

	var json_author_data []interface {}

	json.Unmarshal(author_data, &json_author_data)

	for _, element := range json_author_data {
		detail := element.(map[string]interface {})

		delete(detail, "password")
	}

	return c.RenderJSON(json_author_data)
}

func (c Author) Show() revel.Result {
	if c.Session["username"] == "" {
		var errors []interface {}

		json.Unmarshal([]byte("[{\"Message\": \"請先登入\"}]"), &errors)

		c.Response.Status = 404

		return c.RenderJSON(Errors{errors})
	}

	who := c.Params.Route.Get("author")

	file, _ := ioutil.ReadFile(revel.BasePath + "/storage/authors.json")

	author_data := []byte(file)

	var json_author_data []interface {}

	json.Unmarshal(author_data, &json_author_data)

	for _, element := range json_author_data {
		detail := element.(map[string]interface {})

		if detail["username"] == who {
			delete(detail, "password")

			return c.RenderJSON(detail)
		}
	}

	var errors []interface {}

	json.Unmarshal([]byte("[{\"Message\": \"沒有這個使用者\"}]"), &errors)

	c.Response.Status = 404

	return c.RenderJSON(Errors{errors})
}

func (c Author) Store(input Input, author models.Author) revel.Result {
    /*   
     *   username
     *   password
     *   name
     *   gender
     *   address
     */
	c.Params.BindJSON(&input)

	c.Validation.Required(input.Username).Message("username 必填")
	c.Validation.Required(input.Password).Message("password 必填")
	c.Validation.Required(input.Name).Message("name 必填")
	c.Validation.Required(input.Gender).Message("gender 必填")
	c.Validation.Match(input.Gender, regexp.MustCompile("^[fmo]{1}$")).Message("gender 必須為 f, m, o 其中一項")
	c.Validation.Required(input.Address).Message("address 必填")

	if c.Validation.HasErrors() {
		c.Response.Status = 400

		return c.RenderJSON(c.Validation)
	}

	file, err := ioutil.ReadFile(revel.BasePath + "/storage/authors.json")

	check(err)

	author_data := []byte(file)

	var json_author_data []interface{}

	err1 := json.Unmarshal(author_data, &json_author_data)

	check(err1)

	for _, element := range json_author_data {
		q := element.(map[string]interface {})

		for k, v := range q {
			if k == "username" {
				c.Validation.Required(input.Username != v).Message("username 已被使用")

				if c.Validation.HasErrors() {
					c.Response.Status = 400

					return c.RenderJSON(c.Validation)
				}
			}
		}
	}

	json_author_data = append(json_author_data, input)

	author_data, err2 := json.Marshal(json_author_data)

	check(err2)

	err3 := ioutil.WriteFile(revel.BasePath + "/storage/authors.json", author_data, 0644)

	check(err3)

	c.Response.Status = 201

	return c.RenderJSON(author)
}

func (c Author) Update(input Input) revel.Result {
	if c.Session["username"] == "" {
		var errors []interface{}

		json.Unmarshal([]byte("[{\"Message\": \"請先登入\"}]"), &errors)

		c.Response.Status = 404

		return c.RenderJSON(Errors{errors})
	}

	who := c.Params.Route.Get("author")

	if c.Session["username"] != who {
		var errors []interface{}

		json.Unmarshal([]byte("[{\"Message\": \"不可以亂改別人喔 <3\"}]"), &errors)

		c.Response.Status = 401

		return c.RenderJSON(Errors{errors})
	}

	c.Params.BindJSON(&input)

	file, _ := ioutil.ReadFile(revel.BasePath + "/storage/authors.json")

	author_data := []byte(file)

	var origin_json_author_data []interface{}

	var json_author_data []interface{}

	var response_author_data map[string]interface{}

	json.Unmarshal(author_data, &origin_json_author_data)

	patch := false

	for _, element := range origin_json_author_data {
		detail := element.(map[string]interface{})

		if detail["username"] == who {
			if input.Gender != "" {
				c.Validation.Match(input.Gender, regexp.MustCompile("^[fmo]{1}$")).Message("gender 必須為 f, m, o 其中一項")

				if c.Validation.HasErrors() {
					c.Response.Status = 400

					return c.RenderJSON(c.Validation)
				}

				detail["gender"] = input.Gender
			}

			if input.Address != "" {
				detail["address"] = input.Address
			}

			if input.Name != "" {
				detail["name"] = input.Name
			}

			patch = true

			response_author_data = detail
		}

		json_author_data = append(json_author_data, detail)
	}

	if patch {
		author_data, err := json.Marshal(json_author_data)

		check(err)

		err1 := ioutil.WriteFile(revel.BasePath + "/storage/authors.json", author_data, 0644)

		check(err1)

		return c.RenderJSON(response_author_data)
	} else {
		var errors []interface{}

		json.Unmarshal([]byte("[{\"Message\": \"沒有這個使用者\"}]"), &errors)

		c.Response.Status = 404

		return c.RenderJSON(Errors{errors})
	}
}
