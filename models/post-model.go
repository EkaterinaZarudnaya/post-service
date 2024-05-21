package models

import "time"

type Post struct {
	ID        *int       `json:"id,omitempty" gorm:"primaryKey"`
	UserEmail string     `json:"user_email"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
}
