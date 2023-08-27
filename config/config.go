package config

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config represents the application configuration
type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	RedisAddr  string
	RedisPass  string
}

// LoadConfig loads the configuration from environment variables or a config file
func LoadConfig() *Config {
	return &Config{
		DBUsername: "user",
		DBPassword: "password",
		DBHost:     "127.0.0.1",
		DBPort:     "3310",
		DBName:     "test",
		RedisAddr:  "localhost:6378",
		RedisPass:  "pwd123",
	}

}

// ConnectDB connects to the MySQL database using GORM
func ConnectDB(cfg *Config) (*gorm.DB, error) {
	dsn := cfg.DBUsername + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ConnectRedis connects to the Redis server
func ConnectRedis(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       0,
	})
}
