package repositories

import (
	"article/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db    *gorm.DB
	redis *redis.Client
	c     *gin.Context
}

func NewArticleRepository(db *gorm.DB, redis *redis.Client) *ArticleRepository {
	return &ArticleRepository{db: db, redis: redis}
}

func (r *ArticleRepository) CreateArticle(article *models.Article) error {
	err := r.db.Create(article).Error
	if err != nil {
		return err
	}

	// Cache the new article
	r.cacheArticle(article)

	return nil
}

func (r *ArticleRepository) deleteCacheKeysByPattern(pattern string) error {
	ctx := context.Background()

	var cursor uint64
	for {
		keys, nextCursor, err := r.redis.Scan(ctx, cursor, pattern, 10).Result()
		if err != nil {
			return err
		}

		for _, key := range keys {
			_ = r.redis.Del(ctx, key).Err()
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return nil
}

func (r *ArticleRepository) cacheArticle(article *models.Article) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("article:%d", article.ID)

	// Convert the article to JSON
	articleJSON, _ := json.Marshal(article)

	// Cache the article with an expiration time
	err := r.redis.Set(ctx, cacheKey, articleJSON, 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error caching article:", err)
	}

	// Delete the cached data for the "GET /articles" route
	cacheKeyGET := "/articles:*" // Cache key for the "GET /articles" route
	if err := r.deleteCacheKeysByPattern(cacheKeyGET); err != nil {
		fmt.Println("Failed to delete cache keys:", err)
	}

}

func (r *ArticleRepository) GetArticles(filterAuthor, keyword string) ([]models.Article, error) {
	ctx := context.Background()
	var articles []models.Article

	cacheKey := fmt.Sprintf("/articles:author=%d:keyword=%d", filterAuthor, keyword)
	cachedData, err := r.redis.Get(ctx, cacheKey).Result()

	if err == nil {
		if jsonErr := json.Unmarshal([]byte(cachedData), &articles); jsonErr == nil {
			return articles, nil
		}
	}
	query := r.db.Model(&models.Article{})

	if filterAuthor != "" {
		query = query.Where("author = ?", filterAuthor)
	}

	if keyword != "" {
		query = query.Where("title LIKE ? OR body LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	query = query.Order("created_at desc").Find(&articles)

	articleJSON, _ := json.Marshal(articles)

	err = r.redis.Set(ctx, cacheKey, articleJSON, 5*time.Hour).Err()
	if err != nil {
		fmt.Println("Error caching article:", err)
	}

	return articles, query.Error
}
