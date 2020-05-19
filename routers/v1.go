package routers

import (
	"beew/controllers"
	"beew/filters"

	"github.com/astaxie/beego"
)

func routesForV1() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("token", &controllers.AuthController{}, "post:Token"),
		//need auth
		beego.NSNamespace("user",
			beego.NSBefore(filters.JwtAuth),
			//beego.NSRouter()
		),
		beego.NSNamespace("article",
			//beego.NSRouter(),
			beego.NSNamespace("/",
				//need auth
				beego.NSBefore(filters.JwtAuth),
			//beego.NSRouter()
			),
		),
	)
	beego.AddNamespace(ns)
}
