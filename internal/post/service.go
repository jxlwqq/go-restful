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
	Update()
	Delete()
}

type CreateRequest struct {
	Name string
	Body string
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
	panic("implement me")
}

func (s service) Update() {
	panic("implement me")
}

func (s service) Delete() {
	panic("implement me")
}

func NewService(db *database.DB) Service  {
	return service{db}
}
