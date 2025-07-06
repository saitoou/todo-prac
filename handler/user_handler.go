//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/usecase/user_usecase.go

package handler

import (
	"context"
	"net/http"
	openapi "todo-golang/gen/openapi/common"
	"todo-golang/usecase"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user *usecase.UserUsecaseInput) error
	// FindUserByEmail(ctx context.Context, email string) (*usecase.UserUsecaseOutput, error)
	Login(ctx context.Context, email, password string) (*string, error)
	// UpdateUser(ctx context.Context, user *usecase.TodoUsecaseInput) error
	// ChangePassword(ctx context.Context, user *usecase.UserUsecaseInput) error
	// DeleteUser(ctx context.Context, user *usecase.UserUsecaseInput) error
	Logout(ctx context.Context, email string) error
}

type UserHandler struct {
	userUsecase UserUsecase
}

func NewUserHandler(userUsecase UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) PostAuthSignup(c echo.Context) error {
	ctx := c.Request().Context()
	req := openapi.SignupRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error"})
	}

	singUpUser := usecase.UserUsecaseInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.userUsecase.CreateUser(ctx, &singUpUser); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesaage": "登録処理失敗"})
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *UserHandler) PostAuthLogin(c echo.Context) error {
	ctx := c.Request().Context()
	req := openapi.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error"})
	}

	loginUser := usecase.UserUsecaseInput{
		Email:    req.Email,
		Password: req.Password,
	}

	ret, err := h.userUsecase.Login(ctx, loginUser.Email, loginUser.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "login failed"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": ret,
	})

}

// TODO
func (h *UserHandler) PostAuthLogout(c echo.Context) error {
	return nil
}
