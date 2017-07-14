package controllers

import (
	"github.com/revel/revel"
	"io/ioutil"
	"encoding/json"
)

type App struct {
	*revel.Controller
}

type Session map[string]string

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Errors struct {
	Errors []interface{}
}

//func check(e error) {
//	if e != nil {
//		panic(e)
//	}
//}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) PostLogin(credential Login) revel.Result {
	c.Params.BindJSON(&credential)

	c.Validation.Required(credential.Username).Message("username 必填")
	c.Validation.Required(credential.Password).Message("password 必填")

	if c.Validation.HasErrors() {
		c.Response.Status = 400

		return c.RenderJSON(c.Validation)
	}

	file, _ := ioutil.ReadFile(revel.BasePath + "/storage/authors.json")

	author_data := []byte(file)

	var json_author_data []interface {}

	json.Unmarshal(author_data, &json_author_data)

	for _, element := range json_author_data {
		detail := element.(map[string]interface {})

		if detail["username"] == credential.Username && detail["password"] == credential.Password {
			c.Session["username"] = credential.Username

			delete(detail, "password")

			return c.RenderJSON(detail)
		}
	}

	var errors []interface {}

	json.Unmarshal([]byte("[{\"Message\": \"帳號或密碼錯誤\"}]"), &errors)

	c.Response.Status = 400

	return c.RenderJSON(Errors{errors})
}

func (c App) PostLogout() revel.Result {
	delete(c.Session, "username")

	return c.RenderText("")
}

func (c App) GetLogin() revel.Result {
	if c.Session["username"] != "" {
		file, _ := ioutil.ReadFile(revel.BasePath + "/storage/authors.json")

		author_data := []byte(file)

		var json_author_data []interface {}

		json.Unmarshal(author_data, &json_author_data)

		for _, element := range json_author_data {
			detail := element.(map[string]interface {})

			if detail["username"] == c.Session["username"] {
				delete(detail, "password")

				return c.RenderJSON(detail)
			}
		}
	}

	var errors []interface {}

	json.Unmarshal([]byte("[{\"Message\": \"沒有登入 ^^\"}]"), &errors)

	c.Response.Status = 401

	return c.RenderJSON(Errors{errors})
}