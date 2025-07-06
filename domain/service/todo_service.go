package service

import (
	"context"
	"log"
	"time"
	"todo-golang/domain/entity"
	"todo-golang/domain/repository"
	"todo-golang/utils"
)

type TodoService struct {
	todoRepo repository.TodoRepository
	userRepo repository.UserRepository
}

func NewTodoService(todoRepo repository.TodoRepository, userRepo repository.UserRepository) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
		userRepo: userRepo,
	}
}

// func (svc *TodoService) GetAllTodos(ctx context.Context) ([]*entity.Todo, error) {
// 	ret, err := svc.todoRepo.FindAll(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ret, nil
// }

func (svc *TodoService) GetTodoByID(ctx context.Context, id uint, email string) (*entity.Todo, error) {
	user, err := svc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	ret, err := svc.todoRepo.FindByUserIDAndID(ctx, id, user.ID)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (svc *TodoService) GetTodosByUserID(ctx context.Context, email string) ([]*entity.Todo, error) {

	user, err := svc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	ret, err := svc.todoRepo.FindByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (svc *TodoService) CreateTodo(ctx context.Context, email string, todo *entity.Todo) error {

	user, err := svc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	todo.CreatedAt = time.Now().In(utils.JstLocation())
	todo.UserID = user.ID

	log.Printf("todo :%v", todo)

	err = svc.todoRepo.Create(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}

func (svc *TodoService) UpdateTodo(ctx context.Context, todo *entity.Todo) error {

	todo.UpdatedAt = time.Now().In(utils.JstLocation())

	if err := svc.todoRepo.Update(ctx, todo); err != nil {
		return err
	}

	return nil
}

func (svc *TodoService) DeleteTodo(ctx context.Context, id uint) error {
	if err := svc.todoRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
