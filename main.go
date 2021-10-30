package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

var (
	RowDataNotFound = errors.Errorf("record data not found")
)

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Dao interface {
	Get(id uint64) interface{}
	List() interface{}
	Create()
	Update()
	Delete(id uint64)
}

func (user *User) Get(id uint64) (*User, err) {
	err := db.Where("id = ?", id).Find(user).Error

	if errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, fmt.Sprintf("user nill,user id: %v", id))
	}
	return &user, nil
}

type Service struct{}

func (s *Service) FindUserByID(userID int) (*model.User, error) {
	return dao.FindUserByID(userID)
}
