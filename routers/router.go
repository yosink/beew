package routers

import (
	"beew/controllers"
	"beew/utils/m_cache"
	"fmt"
	"time"

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
		cache, _ := m_cache.GetCache("redis")
		cache.Put("bee_cahce", "bee_cache", 5*time.Minute)
		c.Output.Body([]byte("test"))
	})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/:id", &controllers.ArticleController{}, "put:Put")
	//beego.Router("/cate", &controllers.CategoryController{})
	//ns := beego.NewNamespace("/v1",
	//	beego.NSInclude(&controllers.CategoryController{}),
	//	beego.NSGet("/ca", func(c *context.Context) {
	//		c.Output.Body([]byte("v1 ca"))
	//	}),
	//)
	//beego.AddNamespace(ns)
	//beego.Include(&controllers.CategoryController{})
}
