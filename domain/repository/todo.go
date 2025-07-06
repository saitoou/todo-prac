//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"
	"todo-golang/domain/entity"
)

type TodoRepository interface {
	// FindAll(ctx context.Context) ([]*entity.Todo, error)
	FindByUserIDAndID(ctx context.Context, id uint, userId uint) (*entity.Todo, error)
	FindByUserID(ctx context.Context, userID uint) ([]*entity.Todo, error)
	Create(ctx context.Context, todo *entity.Todo) error
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id uint) error
}
