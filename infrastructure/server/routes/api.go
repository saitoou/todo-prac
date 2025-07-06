package routes

import (
	"todo-golang/infrastructure/container"
	"todo-golang/infrastructure/middleware"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

func registerAPIRoutes(e *echo.Echo, apiContainer *container.APIContainer) {

	api := e.Group("/api")

	api.Use(mw.CORSWithConfig(mw.DefaultCORSConfig))
	api.Use(middleware.JWTMiddleware)

	// 認証無しのエンドポイント（ログイン、サインアップ）を追加
	// api.POST("/auth/login", apiContainer.Todo.PostAuthLogin)
	// api.POST("/auth/signup", apiContainer.Todo.PostAuthSignup)
	// TODO JWT

	registerTodoRoutes(api, apiContainer)
}
