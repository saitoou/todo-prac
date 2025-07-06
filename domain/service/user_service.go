package service

import (
	"context"
	"errors"
	"log"
	"todo-golang/domain/entity"
	"todo-golang/domain/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (svc *UserService) FindUser(ctx context.Context, email, password string) (*entity.User, error) {

	ret, err := svc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid Email error")
	}
	log.Print("find user in service")
	log.Print(password)
	log.Print(ret.Password)

	if !ret.IsValidPassword(password) {
		return nil, errors.New("invalid Password Error")
	}

	log.Print("success valid password")

	return ret, nil
}

func (svc *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	if err := svc.userRepo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}
