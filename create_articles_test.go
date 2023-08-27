package main

import (
	"article/handlers"
	"article/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateArticle(t *testing.T) {
	// Create a mock article service
	service := new(MockArticleService)

	// Create a test handler with the mock service
	handler := handlers.NewArticleHandler(service)

	// Create a test article
	newArticle := models.Article{
		Author: "Test Author",
		Title:  "Test Title",
		Body:   "Test Body",
	}

	// Marshal the article into JSON
	payload, _ := json.Marshal(newArticle)

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Create a test request
	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(payload))

	// Create a Gin context using the recorder and request
	context, _ := gin.CreateTestContext(recorder)
	context.Request = req

	// Mock the service's CreateArticle method
	service.On("CreateArticle", mock.AnythingOfType("*models.Article")).Return(nil)

	// Perform the request
	handler.CreateArticle(context)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// Assert that the service's method was called with the correct argument
	service.AssertExpectations(t)
}
