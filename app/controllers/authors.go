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
	return c.Todo()
}

func (c Author) Show() revel.Result {
	return c.Todo()
}

func (c Author) Store(author models.Author, input Input) revel.Result {
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
	c.Validation.Match(input.Gender, regexp.MustCompile("[fmo]")).Message("gender 必須為 f, m, o 其中一項")
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
		q := element.(map[string]interface{})
		for k, v := range q {
			if k == "username" && v == input.Username {
				c.Response.Status = 400

				var msg []interface{}

				jmsg := []byte(`{"Errors": [{"Message": "username 已存在","Key": "input.Username"}]}`)

				json.Unmarshal(jmsg, &msg)

				return c.RenderJSON(msg)
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

func (c Author) Update() revel.Result {
	return c.Todo()
}
