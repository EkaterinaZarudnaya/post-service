package models

type Post struct {
	Id        *int    `json:"id,omitempty"`
	UserEmail string  `json:"user_email"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	CretedAt  string  `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}
