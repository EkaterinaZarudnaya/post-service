package service

import (
	"errors"
	"post-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostService interface {
	NewPost(post models.Post) error
	GetPosts() ([]models.Post, error)
	GetPostById(id int) (models.Post, error)
	UpdatePostById(id int, post models.Post) error
	DeletePostById(id int) error
}

type Service struct {
	db *gorm.DB
}

func NewService(dsn string) (*Service, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		return nil, err
	}

	return &Service{db: db}, nil
}

func (s *Service) NewPost(post models.Post) error {
	return s.db.Create(&post).Error
}

func (s *Service) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := s.db.Find(&posts).Error
	return posts, err
}

func (s *Service) GetPostById(id int) (models.Post, error) {
	var post models.Post
	err := s.db.First(&post, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return post, errors.New("post not found")
	}
	return post, err
}

func (s *Service) UpdatePostById(id int, updatedPost models.Post) error {
	var existingPost models.Post
	err := s.db.First(&existingPost, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("post not found")
	}
	
	updatedPost.ID = existingPost.ID
	return s.db.Model(&existingPost).Updates(updatedPost).Error
}

func (s *Service) DeletePostById(id int) error {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return err
	}
	
	err := s.db.Delete(&post).Error
	return err
}
