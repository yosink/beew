package routers

import (
	"beew/controllers"
	"beew/filters"
	"beew/utils"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Get("/auth", func(c *context.Context) {
		url := "http://www.baidu.com"
		res := govalidator.IsURL(url)
		fmt.Println(res)
		c.Output.Body([]byte("url checked"))
	})
	beego.Get("/test", func(c *context.Context) {
		token, err := utils.GenerateJwtToken(1)
		if err != nil {
			fmt.Printf("%v", err)
		}
		c.Output.JSON(token, false, false)
	})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/:id", &controllers.ArticleController{}, "put:Put")

	routesForV1()

	beego.InsertFilter("/auth", beego.BeforeRouter, filters.JwtAuth)
}
