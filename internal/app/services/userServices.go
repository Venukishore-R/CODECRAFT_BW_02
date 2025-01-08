package services

import (
	"fmt"
	"regexp"

	"github.com/Venukishore-R/CODECRAFT_BW_02/internal/app/models"
	"github.com/google/uuid"
)

func (s *Service) Create(user *models.User) error {
	if !s.IsValidEmail(user.Email) {
		return fmt.Errorf("invalid email: %v", user.Email)
	}

	user.Id = s.CreateUUID()

	return s.UserModel.Create(user)
}

func (s *Service) FindAll() ([]*models.User, error) {
	return s.UserModel.FindAll()
}

func (s *Service) FindByEmail(email string) (*models.User, error) {
	return s.UserModel.FindByEmail(email)
}

func (s *Service) Update(email string, user *models.User) error {
	return s.UserModel.Update(email, user)
}

func (s *Service) Delete(email string) error {
	return s.UserModel.Delete(email)
}

func (s *Service) CreateUUID() string {
	id := uuid.New()
	return id.String()
}

func (s *Service) IsValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	reg := regexp.MustCompile(emailPattern)

	return reg.MatchString(email)
}
