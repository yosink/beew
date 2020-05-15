package controllers

import (
	"beew/repositories"
	"beew/services"
	"beew/utils"
	"beew/utils/my_logger"
	"beew/validators"
	"net/http"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
	service services.IArticleService
}

func (c *ArticleController) Prepare() {
	repo := repositories.NewArticleRepo()
	c.service = services.NewArticleService(repo)
}

func (c *ArticleController) Get() {
	var (
		app      = utils.B{C: c.Ctx}
		httpCode = http.StatusOK
		errCode  = 200
		message  string
		data     interface{}
	)

	list, err := c.service.GetList(1, 15, map[string]interface{}{})
	if err != nil {
		httpCode = http.StatusInternalServerError
		errCode = 500
		message = err.Error()
	} else {
		data = list
	}
	app.JsonResponse(httpCode, errCode, message, data)
}

func (c *ArticleController) Post() {
	var (
		app      = utils.B{C: c.Ctx}
		httpCode = http.StatusOK
		errCode  = 200
		message  string
		data     interface{}
		article  validators.ArticleAdd
	)
	err := c.ParseForm(&article)
	if err != nil {
		httpCode = 400
		errCode = 400
		message = "绑定form失败"
	} else {
		if err = validators.BindAndValidate(article); err == nil {
			// create
			insertMap := map[string]interface{}{
				"category_id":      article.CategoryID,
				"user_id":          article.UserID,
				"slug":             article.Slug,
				"title":            article.Title,
				"subtitle":         article.Subtitle,
				"content":          article.Content,
				"page_image":       article.PageImage,
				"meta_description": article.MetaDescription,
				"recommend":        article.Recommend,
				"sort":             article.Sort,
				"state":            article.State,
				"view_count":       article.ViewCount,
			}
			_, err := c.service.Add(insertMap)
			if err != nil {
				errCode = 4001
				message = "创建文章失败"
				my_logger.Error("c.service.Add error: %v", err)
			}
		} else {
			httpCode = 422
			errCode = 422
			message = err.Error()
		}
	}
	app.JsonResponse(httpCode, errCode, message, data)

}