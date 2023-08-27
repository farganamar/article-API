package services

import (
	"article/models"
	"article/repositories"
)

type ArticleService interface {
	CreateArticle(article *models.Article) error
	GetArticles(filterAuthor, keyword string) ([]models.Article, error)
}

type ArticleServiceImpl struct {
	repo *repositories.ArticleRepository
}

func NewArticleService(repo *repositories.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{repo: repo}
}

func (s *ArticleServiceImpl) CreateArticle(article *models.Article) error {
	err := s.repo.CreateArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func (s *ArticleServiceImpl) GetArticles(filterAuthor, keyword string) ([]models.Article, error) {
	return s.repo.GetArticles(filterAuthor, keyword)
}
