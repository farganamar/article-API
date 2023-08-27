package main

import (
	"article/handlers"
	"article/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetArticles(t *testing.T) {
	// Create a mock article service
	service := new(MockArticleService)

	// Create a test handler with the mock service
	handler := handlers.NewArticleHandler(service)

	// Create a mock gin context
	context := &gin.Context{}

	// Mock the service's GetArticles method
	expectedArticles := []models.Article{
		{ID: 1, Author: "Author 1", Title: "Title 1", Body: "Body 1"},
		{ID: 2, Author: "Author 2", Title: "Title 2", Body: "Body 2"},
	}
	service.On("GetArticles", "", "").Return(expectedArticles, nil)

	// Set the request
	context.Request, _ = http.NewRequest("GET", "/articles", nil)

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Set the context's Writer to the recorder
	context.Writer = recorder

	// Perform the request
	handler.GetArticles(context)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Unmarshal the response body
	var responseArticles []models.Article
	err := json.Unmarshal(recorder.Body.Bytes(), &responseArticles)
	assert.NoError(t, err)

	// Assert that the response body matches the expected articles
	assert.Equal(t, expectedArticles, responseArticles)

	// Assert that the service's method was called
	service.AssertExpectations(t)
}
