package apis

import (
	"beego-app/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) List() {
	users := models.GetUsers()
	//	c.Data["users"] = users

	c.Data["json"] = users
	c.ServeJSON()

}
