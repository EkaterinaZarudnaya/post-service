package handlers

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/ekaterinazarudnaya/post-service/models"
	service "github.com/ekaterinazarudnaya/post-service/services"

	_ "github.com/ekaterinazarudnaya/post-service/docs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  service.PostService
	validate *validator.Validate
}

func NewPostHandler(service service.PostService) *Handler {
	return &Handler{
		service:  service,
		validate: validator.New(),
	}
}

// NewPost             godoc
// Handler functions
// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param   post body models.Post true "Post to create"
// @Success 200 {object} models.Post
// @Router /api/v1/posts/ [post]
func (h *Handler) NewPost(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON input"})
		return
	}

	if err := h.validate.Struct(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	if err := h.service.NewPost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPosts             godoc
// @Summary Get all posts
// @Description Get all posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Post
// @Router /api/v1/posts/ [get]
func (h *Handler) GetPosts(c *gin.Context) {
	posts, err := h.service.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPostById             godoc
// @Summary Get post by ID
// @Description Get post by ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Param   id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /api/v1/posts/{id} [get]
func (h *Handler) GetPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	post, err := h.service.GetPostById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePostById             godoc
// @Summary Update post by ID
// @Description Update post by ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Param   id path int true "Post ID"
// @Param   post body models.Post true "Post to update"
// @Success 200 {object} models.Post
// @Router /api/v1/posts/{id} [put]
func (h *Handler) UpdatePostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON input"})
		return
	}

	if err := h.validate.Struct(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	if err := h.service.UpdatePostById(id, post); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	newPost, err := h.service.GetPostById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, newPost)
}

// DeletePostById             godoc
// @Summary Delete post by ID
// @Description Delete post by ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Param   id path int true "Post ID"
// @Success 204
// @Router /api/v1/posts/{id} [delete]
func (h *Handler) DeletePostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	if err := h.service.DeletePostById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.Status(http.StatusOK)
}
