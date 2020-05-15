package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	data := map[string]interface{}{"website":"beego.me"}
	c.Data["json"] = data
	c.ServeJSON()

}
