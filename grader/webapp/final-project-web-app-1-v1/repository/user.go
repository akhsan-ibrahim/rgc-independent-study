package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var result entity.User
	err := r.db.WithContext(ctx).Table("users").Select("*").Where("id = ?", id).Find(&result).Error
	return result, err // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var result entity.User
	err := r.db.WithContext(ctx).Table("users").Select("*").Where("email = ?", email).Find(&result).Error
	return result, err // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	data, _ := r.GetUserByID(ctx, user.ID)
	// data := entity.User{}
	// r.db.Model(entity.User{}).Where("id = ?", user.ID).Find(&data)
	return data, err // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	var result entity.User
	err := r.db.WithContext(ctx).Table("users").Where("id = ?", user.ID).Updates(user).Error
	return result, err // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.User{}, id).Error // TODO: replace this
}
