package controllers

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

type Post struct {
	Title    string `valid:"alphanum,required"`
	Message  string `valid:"duck,ascii"`
	Message2 string `valid:"animal(dog)"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}

// @router /c_test
func (ctl *CategoryController) MyTest() {
	ctl.Ctx.Output.Body([]byte("my test"))
}

//func (ctl *CategoryController) Get() {
//	post := &Post{
//		Title:    "My Example Post",
//		Message:  "duck",
//		Message2: "dog",
//		AuthorIP: "123.234.54.3",
//	}
//
//	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
//		return str == "duck"
//	})
//
//	govalidator.ParamTagMap["animal"] = govalidator.ParamValidator(func(str string, params ...string) bool {
//		species := params[0]
//		return str == species
//	})
//
//	govalidator.ParamTagRegexMap["animal"] = regexp.MustCompile("^animal\\((\\w+)\\)$")
//
//	valid, err := govalidator.ValidateStruct(post)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(valid)
//}
