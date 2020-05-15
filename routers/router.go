package routers

import (
	"beew/controllers"
	"beew/utils/m_cache"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Get("/auth", func(c *context.Context) {
		c.Output.Body([]byte("Auth center"))
	})
	beego.Get("/test", func(c *context.Context) {
		cache, _ := m_cache.GetCache("redis")
		cache.Put("bee_cahce", "bee_cache", 5*time.Minute)
		c.Output.Body([]byte("test"))
	})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/article", &controllers.ArticleController{})
}
