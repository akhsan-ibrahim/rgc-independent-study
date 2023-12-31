package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) AddUser(user model.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *UserRepository) UserAvail(cred model.User) error {
	result := model.User{}
	err := u.db.Table("users").Select("*").Where("username = ? AND password = ?", cred.Username, cred.Password).Find(&result).Error
	if err != nil {
		return err
	}
	if result.ID == 0 {
		return fmt.Errorf("not found")
	}
	return nil // TODO: replace this
}

func (u *UserRepository) CheckPassLength(pass string) bool {
	if len(pass) <= 5 {
		return true
	}

	return false
}

func (u *UserRepository) CheckPassAlphabet(pass string) bool {
	for _, charVariable := range pass {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}
