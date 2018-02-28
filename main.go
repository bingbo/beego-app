package main

import (
	_ "beego-app/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/down", "file")
	beego.Run()
}
