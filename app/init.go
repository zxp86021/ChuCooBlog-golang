package app

import (
	"github.com/revel/revel"
    "io/ioutil"
    "os"
	//"encoding/json"
	//"github.com/zxp86021/ChuCooBlog-golang/app/controllers"
)

type Session map[string]string

type Errors struct {
	Errors []interface{}
}

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}


	// register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	revel.OnAppStart(CheckStorage)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)

	//revel.InterceptFunc(CheckLogin, revel.BEFORE, &controllers.App{})
	//revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
}

// HeaderFilter adds common security headers
// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

func CheckStorage() {
    if _, err := os.Stat(revel.BasePath + "/storage/posts.json"); os.IsNotExist(err) {
        // storage/posts.json does not exist
        err := ioutil.WriteFile(revel.BasePath + "/storage/posts.json", []byte("[]"), 0644)

        if err != nil {
            panic(err)
        }
    }

    if _, err := os.Stat(revel.BasePath + "/storage/authors.json"); os.IsNotExist(err) {
        // storage/authors.json does not exist
        err := ioutil.WriteFile(revel.BasePath + "/storage/authors.json", []byte("[{\"username\":\"test1\",\"password\":\"test123\",\"name\":\"test\",\"gender\":\"f\",\"address\":\"earth\"}]"), 0644)

        if err != nil {
            panic(err)
        }
    }
}

//func CheckLogin(c *revel.Controller) revel.Result {
//	if c.Session["username"] == "" {
//		var errors []interface {}
//
//		json.Unmarshal([]byte("[{\"Message\": \"請先登入\"}]"), &errors)
//
//		c.Response.Status = 401
//
//		return c.RenderJSON(Errors{errors})
//	}
//
//	return nil
//}


//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
