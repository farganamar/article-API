package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func CacheMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a cache key based on the request URL
		cacheKey := c.Request.URL.String()

		// Check if data exists in cache
		ctx := context.Background()
		cachedData, err := redisClient.Get(ctx, cacheKey).Result()

		if err == nil {
			// Parse cached JSON response
			var responseData interface{} // Adjust this to match your response structure
			if err := json.Unmarshal([]byte(cachedData), &responseData); err != nil {
				// Handle unmarshaling error
				c.String(500, "Internal Server Error")
				return
			}

			// Return cached JSON response with proper Content-Type header
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.JSON(200, responseData) // Correctly send the JSON response
			return
		}

		// Capture the response
		buf := &bytes.Buffer{}
		cWriter := &responseWriter{ResponseWriter: c.Writer, Body: buf}
		c.Writer = cWriter

		// Continue processing the request
		c.Next()

		// Cache the response if status is 200 OK
		if c.Writer.Status() == 200 {
			response := buf.String()
			_ = redisClient.Set(ctx, cacheKey, response, 5*time.Hour).Err()
		}
	}
}

type responseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w *responseWriter) Write(data []byte) (int, error) {
	// Store the response body in the Body buffer
	w.Body.Write(data)
	return w.ResponseWriter.Write(data)
}
