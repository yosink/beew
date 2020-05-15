package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get()  {
	c.Data["json"] = map[string]interface{} {
		"code":200,
		"message":"success",
	}
	c.ServeJSON()
}
