package routers

import (
	"beego-app/apis"
	"beego-app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/index", &controllers.UserController{}, "*:Index")
	beego.Router("/user/list", &controllers.UserController{}, "get:List")

	ns := beego.NewNamespace("/api",
		beego.NSRouter("/user/list", &apis.UserController{}, "get:List"),
	)
	beego.AddNamespace(ns)
}
