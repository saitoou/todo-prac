package repository

import (
	"context"
	"todo-golang/domain/entity"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// func (r *TodoRepository) FindAll(ctx context.Context) ([]*entity.Todo, error) {
// 	var todos []*entity.Todo
// 	if err := r.db.Find(&todos).Error; err != nil {
// 		return nil, err
// 	}
// 	return todos, nil
// }

func (r *TodoRepository) FindByUserIDAndID(ctx context.Context, id uint, userId uint) (*entity.Todo, error) {
	var todo *entity.Todo
	if err := r.db.First(&todo, id, userId).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) FindByUserID(ctx context.Context, userID uint) ([]*entity.Todo, error) {
	var todos []*entity.Todo
	if err := r.db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	if err := r.db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) Update(ctx context.Context, todo *entity.Todo) error {
	if err := r.db.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.Delete(&entity.Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}
