package repositories

import (
	"beew/datasource"
	"beew/models"
	"beew/utils"
	"beew/utils/formater"
	"fmt"

	"github.com/jinzhu/gorm"
)

type IArticleRepo interface {
	GetList(int, int, map[string]interface{}) ([]*models.Article, error)
	GetByID(int) (models.Article, error)
	Add(map[string]interface{}) (bool, error)
	Edit(int, map[string]interface{}) (bool, error)
	Remove(int) error
}

type ArticleRepo struct {
	DB *gorm.DB
}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{DB: datasource.NewDbInstance()}
}

func (a *ArticleRepo) GetList(page int, perPage int, query map[string]interface{}) (result []*models.Article, err error) {
	err = a.DB.Where(query).Offset(utils.GetPageNum(page, perPage)).Limit(perPage).Find(&result).Error
	return
}

func (a *ArticleRepo) GetByID(id int) (result models.Article, err error) {
	err = a.DB.Where("id = ?", id).Find(&result).Error
	return
}

func (a *ArticleRepo) Add(data map[string]interface{}) (bool, error) {
	article := models.Article{
		CategoryID:      data["category_id"].(int64),
		UserID:          data["user_id"].(int64),
		Slug:            data["slug"].(string),
		Title:           data["title"].(string),
		Subtitle:        data["subtitle"].(string),
		Content:         data["content"].(string),
		PageImage:       data["page_image"].(string),
		MetaDescription: data["meta_description"].(string),
		Recommend:       data["recommend"].(int8),
		Sort:            data["sort"].(int),
		State:           0,
		ViewCount:       data["view_count"].(int),
		PublishedAt:     formater.XTime{},
	}

	err := a.DB.Create(&article).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArticleRepo) Edit(id int, m map[string]interface{}) (bool, error) {
	var article models.Article
	if err := a.DB.Where("id = ?", id).First(&article).Error; gorm.IsRecordNotFoundError(err) {
		return false, fmt.Errorf("article is not found")
	}
	err := a.DB.Model(&models.Article{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArticleRepo) Remove(id int) error {
	return a.DB.Where("id = ?", id).Delete(&ArticleRepo{}).Error
}
