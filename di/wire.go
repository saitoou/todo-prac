//go:build wireinject
// +build wireinject

package di

import (
	"todo-golang/domain/repository"
	"todo-golang/domain/service"
	"todo-golang/gen/openapi/v1"
	"todo-golang/handler"
	"todo-golang/infrastructure/container"
	repo "todo-golang/infrastructure/repository"
	"todo-golang/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var repositorySet = wire.NewSet(
	repo.NewTodoRepository,
	repo.NewUserRepository,
	wire.Bind(new(repository.TodoRepository), new(*repo.TodoRepository)),
	wire.Bind(new(repository.UserRepository), new(*repo.UserRepository)),
)

var serviceSet = wire.NewSet(
	service.NewTodoService,
	service.NewUserService,
	wire.Bind(new(usecase.TodoService), new(*service.TodoService)),
	wire.Bind(new(usecase.UserService), new(*service.UserService)),
)

var usecaseSet = wire.NewSet(
	usecase.NewTodoUseCase,
	usecase.NewUserUsecase,
	wire.Bind(new(handler.TodoUseCase), new(*usecase.TodoUsecase)),
	wire.Bind(new(handler.UserUsecase), new(*usecase.UserUsecase)),
)

type handlerWrapper struct {
	*handler.TodoHandler
	*handler.UserHandler
}

var handlerSet = wire.NewSet(
	handler.NewTodoHandler,
	handler.NewUserHandler,
	wire.Struct(new(handlerWrapper), "*"),
	wire.Bind(new(openapi.ServerInterface), new(*handlerWrapper)),
)

// InitializeAPIContainer function now builds everything together using wire.Build
func InitializeAPIContainer(
	db *gorm.DB,
) *container.APIContainer {
	wire.Build(
		repositorySet,
		serviceSet,
		usecaseSet,
		handlerSet,
		wire.Struct(new(container.APIContainer), "*"), // Container struct will hold the dependencies
	)
	return &container.APIContainer{} // Returning the container that holds all the injected dependencies
}
