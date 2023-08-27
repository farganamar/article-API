package main

import (
	"article/config"
	"article/handlers"
	"article/migrations"
	"article/repositories"
	"article/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MySQL
	db, err := config.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}
	migrations.Migrate(db)

	// Connect to Redis
	redisClient := config.ConnectRedis(cfg)

	// Initialize repositories
	articleRepo := repositories.NewArticleRepository(db, redisClient)

	// Initialize services
	articleService := services.NewArticleService(articleRepo)

	// Initialize handlers
	articleHandler := handlers.NewArticleHandler(articleService)

	// Create the Gin engine
	r := gin.Default()

	// Use cache middleware
	// r.Use(middleware.CacheMiddleware(redisClient))

	// Define routes
	r.POST("/articles", articleHandler.CreateArticle)
	r.GET("/articles", articleHandler.GetArticles)

	// Start the server
	r.Run(":8080")
}
