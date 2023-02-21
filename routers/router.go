package routers

import (
	"scylla/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/signup", &controllers.SignupController{}, "get:View;post:Signup")
	beego.Router("/login", &controllers.LoginController{}, "get:View;post:Login")

	ns :=
		beego.NewNamespace("/home",
			beego.NSBefore(Auth),

			beego.NSRouter("/", &controllers.HomeController{}, "get:View"),
			beego.NSRouter("/upload", &controllers.HomeController{}, "post:UploadFile"),
		)

	// register namespace
	beego.AddNamespace(ns)
}
