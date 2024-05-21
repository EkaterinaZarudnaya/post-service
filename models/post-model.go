package models

import "time"

type Post struct {
	ID        *int       `json:"id,omitempty" gorm:"primaryKey"`
	UserEmail string     `json:"user_email" validate:"required,email"`
	Title     string     `json:"title" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
}
