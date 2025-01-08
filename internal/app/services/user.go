package services

import "github.com/Venukishore-R/CODECRAFT_BW_02/internal/app/models"

type Service struct {
	UserModel *models.UserModel
}

func NewService(userModel *models.UserModel) *Service {
	return &Service{
		UserModel: userModel,
	}
}

type UserService interface {
	Create(user *models.User) error
	FindAll() ([]*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(email string, user *models.User) error
	Delete(email string) error
	IsValidEmail(email string) bool
	CreateUUID() string
}
