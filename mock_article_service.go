package main

import (
	"article/models"

	"github.com/stretchr/testify/mock"
)

// MockArticleService is a mock implementation of the ArticleService interface
type MockArticleService struct {
	mock.Mock
}

func (m *MockArticleService) CreateArticle(article *models.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleService) GetArticles(filterAuthor, keyword string) ([]models.Article, error) {
	args := m.Called(filterAuthor, keyword)
	return args.Get(0).([]models.Article), args.Error(1)
}
