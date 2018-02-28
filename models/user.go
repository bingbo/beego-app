package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `json:"id"` //json输出字段别名
	Name     string `json:"name"`
	Password string `json:"password`
	Email    string `json:"email"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	orm.RegisterModel(new(User))
}

func GetUsers() (users []User) {
	o := orm.NewOrm()
	o.Using("default")
	num, err := o.QueryTable("user").All(&users)
	fmt.Println(num, err, users)
	for _, user := range users {
		fmt.Println(user.Id, user.Name, user.Password, user.Email)
	}
	return

}
