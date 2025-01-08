package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func (m *UserModel) Create(user *User) error {
	if err := m.CheckUserExists(user.Email); err != nil {
		return err
	}

	if err := m.db.Model(&User{}).Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (m *UserModel) FindAll() ([]*User, error) {
	var users []*User
	if err := m.db.Model(&User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserModel) FindByEmail(email string) (*User, error) {
	var user *User
	if err := m.db.Model(&User{}).Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.Id == "" {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (m *UserModel) Update(email string, user *User) error {
	existingUser, err := m.FindByEmail(email)
	if err != nil {
		return err
	}

	if existingUser.Email != email {
		return fmt.Errorf("user not found")
	}

	if err := m.db.Model(&User{}).Where("email = ?", email).UpdateColumns(user).Error; err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Delete(email string) error {
	existingUser, err := m.FindByEmail(email)
	if err != nil {
		return err
	}

	if existingUser.Email != email {
		return fmt.Errorf("user not found")
	}

	if err := m.db.Model(&User{}).Where("email = ?", email).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}

func (m *UserModel) CheckUserExists(email string) error {
	var user *User
	err := m.db.Model(&User{}).Where("email = ?", email).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("failed to check user existence", err)
		return fmt.Errorf("failed to check user existence: %v", err)
	}

	log.Println("user", user.Email)
	if user.Id != "" {
		return fmt.Errorf("user already exists")
	}

	return nil
}
