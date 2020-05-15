package main

import (
	"beew/filters"
	_ "beew/routers"
	"beew/utils/my_logger"

	"github.com/astaxie/beego"
)

func main() {
	my_logger.Setup()
	beego.InsertFilter("/auth", beego.BeforeRouter, filters.JwtAuth)
	beego.Run()
}
