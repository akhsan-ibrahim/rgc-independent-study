package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return TodoRepository{db}
}

func (u *TodoRepository) AddTodo(todo model.Todo) error {
	err := u.db.Create(&todo).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *TodoRepository) ReadTodo() ([]model.Todo, error) {
	result := []model.Todo{}
	rows, err := u.db.Table("todos").Select("*").Where("deleted_at IS NULL").Rows()
	if err != nil {
		return []model.Todo{}, err
	}

	defer rows.Close()

	for rows.Next() {
		u.db.ScanRows(rows, &result)
	}
	return result, nil // TODO: replace this
}

func (u *TodoRepository) UpdateDone(id uint, status bool) error {
	err := u.db.Table("todos").Where("id = ?", id).Update("done", status).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *TodoRepository) DeleteTodo(id uint) error {
	err := u.db.Delete([]model.Todo{}, id).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
