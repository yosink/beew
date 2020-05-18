package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

// @router /test_test
func (r *TestController) RouteTest() {
	r.Ctx.Output.Body([]byte("testtest"))
}
