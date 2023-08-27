package models

import "time"

type Article struct {
	ID        uint `gorm:"primaryKey"`
	Author    string
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
