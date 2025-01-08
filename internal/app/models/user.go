package models

import (
	"gorm.io/gorm"
)

type User struct {
	Id    string
	Name  string
	Email string
	Age   int
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

type UserModelRepo interface {
	Create(user *User) error
	FindAll() ([]*User, error)
	FindByEmail(email string) (*User, error)
	Update(email string, user *User) error
	Delete(email string) error
	CheckUserExists(email string) error
}
