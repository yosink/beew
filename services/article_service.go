package services

import (
	"beew/models"
	"beew/repositories"
)

type IArticleService interface {
	GetList(int, int, map[string]interface{}) ([]*models.Article, error)
	GetByID(int) (models.Article, error)
	Add(map[string]interface{}) (bool, error)
	Edit(int, map[string]interface{}) (bool, error)
	Remove(int) error
}

type ArticleService struct {
	repo repositories.IArticleRepo
}

func NewArticleService(repo repositories.IArticleRepo) IArticleService {
	return &ArticleService{repo: repo}
}

func (a *ArticleService) GetList(page int, perPage int, query map[string]interface{}) ([]*models.Article, error) {
	return a.repo.GetList(page, perPage, query)
}

func (a *ArticleService) GetByID(id int) (models.Article, error) {
	return a.repo.GetByID(id)
}

func (a *ArticleService) Add(data map[string]interface{}) (bool, error) {
	return a.repo.Add(data)
}

func (a *ArticleService) Edit(id int, data map[string]interface{}) (bool, error) {
	return a.repo.Edit(id, data)
}

func (a *ArticleService) Remove(id int) error {
	return a.repo.Remove(id)
}
