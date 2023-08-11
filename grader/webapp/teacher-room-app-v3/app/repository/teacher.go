package handler

import (
	"a21hc3NpZ25tZW50/app/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (u *TeacherRepo) AddTeacher(teacher model.Teacher) error {
	return u.db.Create(&teacher).Error
}

func (u *TeacherRepo) ReadTeacher() ([]model.ViewTeacher, error) {
	var result []model.ViewTeacher
	err := u.db.Table("teachers").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *TeacherRepo) UpdateName(id uint, name string) error {
	return u.db.Table("teachers").Where("id = ?", id).Update("name", name).Error
}

func (u *TeacherRepo) DeleteTeacher(id uint) error {
	return u.db.Delete(&model.Teacher{}, id).Error
}
