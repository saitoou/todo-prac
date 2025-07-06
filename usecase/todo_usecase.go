//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/service/todo_service.go

package usecase

import (
	"context"
	"log"
	"time"
	"todo-golang/domain/entity"

	"github.com/samber/lo"
)

type TodoService interface {
	// GetAllTodos(ctx context.Context) ([]*entity.Todo, error)
	GetTodoByID(ctx context.Context, id uint, email string) (*entity.Todo, error)
	GetTodosByUserID(ctx context.Context, email string) ([]*entity.Todo, error)
	CreateTodo(ctx context.Context, email string, todo *entity.Todo) error
	UpdateTodo(ctx context.Context, todo *entity.Todo) error
	DeleteTodo(ctx context.Context, id uint) error
}

type TodoUsecaseInput struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type TodoUsecaseOutput struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type TodoUsecase struct {
	todoSvc TodoService
}

func NewTodoUseCase(todoSvc TodoService) *TodoUsecase {
	return &TodoUsecase{todoSvc: todoSvc}
}

// func (uc *TodoUsecase) GetAllTodos(ctx context.Context, email string) ([]TodoUsecaseOutput, error) {
// 	ret, err := uc.todoSvc.GetAllTodos(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	todos := lo.Map(ret, func(todo *entity.Todo, _ int) TodoUsecaseOutput {
// 		return TodoUsecaseOutput{
// 			Title:   todo.Title,
// 			Content: todo.Content,
// 		}
// 	})

// 	return todos, nil

// }

func (uc *TodoUsecase) GetTodoByID(ctx context.Context, id uint, email string) (*TodoUsecaseOutput, error) {
	ret, err := uc.todoSvc.GetTodoByID(ctx, id, email)
	if err != nil {
		return nil, err
	}

	todo := &TodoUsecaseOutput{
		Title:   ret.Title,
		Content: ret.Content,
	}

	return todo, nil
}

func (uc *TodoUsecase) GetTodosByUserID(ctx context.Context, email string) ([]TodoUsecaseOutput, error) {
	ret, err := uc.todoSvc.GetTodosByUserID(ctx, email)
	if err != nil {
		return nil, err
	}

	todos := lo.Map(ret, func(todo *entity.Todo, _ int) TodoUsecaseOutput {
		return TodoUsecaseOutput{
			Title:   todo.Title,
			Content: todo.Content,
		}
	})

	return todos, nil

}

func (uc *TodoUsecase) CreateTodo(ctx context.Context, email string, todo TodoUsecaseInput) error {

	insertTodo := &entity.Todo{
		Title:   todo.Title,
		Content: todo.Content,
	}
	log.Printf("email :%v", email)
	log.Printf("todo :%+v", todo)

	if err := uc.todoSvc.CreateTodo(ctx, email, insertTodo); err != nil {
		return err
	}

	return nil
}

func (uc *TodoUsecase) UpdateTodo(ctx context.Context, todo TodoUsecaseInput) error {

	updateTodo := &entity.Todo{
		ID:      todo.ID,
		Title:   todo.Title,
		Content: todo.Content,
		// UserID:    todo.UserID,
		// CreatedAt: todo.CreatedAt,
	}

	if err := uc.todoSvc.UpdateTodo(ctx, updateTodo); err != nil {
		return err
	}

	return nil
}

func (uc *TodoUsecase) DeleteTodo(ctx context.Context, id uint) error {

	if err := uc.todoSvc.DeleteTodo(ctx, id); err != nil {
		return err
	}

	return nil
}
