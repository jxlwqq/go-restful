package post

import (
	"github.com/jxlwqq/go-restful/internal/entity"
	"github.com/jxlwqq/go-restful/pkg/database"
)

type Service interface {
	Get(id string) (entity.Post, error)
	Count() (int, error)
	Query(offset, limit int) ([]entity.Post, error)
	Create(req CreateRequest) (entity.Post, error)
	Update(id string, req UpdateRequest) (entity.Post, error)
	Delete(id string) (entity.Post, error)
}

type CreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type UpdateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (r CreateRequest) Validate() error {
	// todo
	return nil
}

type service struct {
	db *database.DB
}

func (s service) Get(id string) (post entity.Post, err error) {
	err = s.db.First(&post, id).Error
	return
}

func (s service) Count() (int, error) {
	panic("implement me")
}

func (s service) Query(offset, limit int) ([]entity.Post, error) {
	panic("implement me")
}

func (s service) Create(req CreateRequest) (post entity.Post, err error) {
	post = entity.Post{
		Title: req.Title,
		Body:  req.Body,
	}
	err = s.db.Create(&post).Error
	return
}

func (s service) Update(id string, req UpdateRequest) (post entity.Post, err error) {
	post, err = s.Get(id)
	if err != nil {
		return
	}
	post.Title = req.Title
	post.Body = req.Body
	s.db.Save(&post)
	return
}

func (s service) Delete(id string) (post entity.Post, err error) {
	post, err = s.Get(id)
	if err != nil {
		return
	}
	s.db.Delete(&post)
	return
}

func NewService(db *database.DB) Service {
	return service{db}
}
