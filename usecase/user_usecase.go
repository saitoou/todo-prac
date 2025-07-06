//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/service/user_service.go

package usecase

import (
	"context"
	"errors"
	"time"
	"todo-golang/config"
	"todo-golang/domain/entity"

	"github.com/golang-jwt/jwt/v5"
)

type UserUsecaseInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserUsecaseOutput struct {
	// ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

type UserService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	// FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	// UpdateUser(ctx context.Context, user *entity.User) error
	// ChangePassword(ctx context.Context, user *entity.User) error
	// DeleteUser(ctx context.Context, user *entity.User) error
	FindUser(ctx context.Context, email, password string) (*entity.User, error)
}

type UserUsecase struct {
	userSvc UserService
}

func NewUserUsecase(userSvc UserService) *UserUsecase {
	return &UserUsecase{userSvc: userSvc}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, input *UserUsecaseInput) error {

	user := entity.NewUser(input.Name, input.Email, input.Password)

	if err := uc.userSvc.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (*string, error) {

	user, err := u.userSvc.FindUser(ctx, email, password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	userOutput := UserUsecaseOutput{
		Name:  user.Name,
		Email: user.Email,
	}

	token, err := generateJWTToken(&userOutput)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// generateToken
func generateJWTToken(user *UserUsecaseOutput) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.Email,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	secret := config.AppConf.Auth.JWT_SECRET

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (u *UserUsecase) Logout(ctx context.Context, email string) error {
	return nil
}

// func (uc *UserUsecase) UpdateUser(ctx context.Context, user UserUsecaseInput) error {
// 	input := &entity.User{
// 		Email:     user.Email,
// 		Name:      user.Name,
// 		CreatedAt: time.Now().In(utils.JstLocation()),
// 	}

// 	if err := uc.userSvc.UpdateUser(ctx, input); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (uc *UserUsecase) ChangePassword(ctx context.Context, user UserUsecaseInput) error {
// 	input := &entity.User{
// 		Email:     user.Email,
// 		Password:  user.Password,
// 		UpdatedAt: time.Now().In(utils.JstLocation()),
// 	}

// 	if err := uc.userSvc.ChangePassword(ctx, input); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (uc *UserUsecase) DeleteUser(ctx context.Context, user UserUsecaseInput) error {
// 	input := &entity.User{
// 		Email: user.Email,
// 		Name:  user.Name,
// 	}

// 	if err := uc.userSvc.DeleteUser(ctx, input); err != nil {
// 		return err
// 	}
// 	return nil
// }
