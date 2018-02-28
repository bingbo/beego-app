package controllers

import (
	"beego-app/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user/index.tpl"
}

func (c *UserController) Index() {
	c.Data["Website"] = "user index page"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user/index.tpl"
}

func (c *UserController) List() {
	users := models.GetUsers()
	c.Data["users"] = users
	c.TplName = "user/list.tpl"

}
