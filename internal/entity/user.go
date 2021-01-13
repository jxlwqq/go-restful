package entity

import (
	"strconv"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (u User) GetID() string {
	return strconv.Itoa(u.ID)
}

func (u User) GetName() string {
	return u.Name
}
