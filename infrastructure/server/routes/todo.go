package routes

import (
	"todo-golang/gen/openapi/v1"
	"todo-golang/infrastructure/container"

	"github.com/labstack/echo/v4"
)

func registerTodoRoutes(g *echo.Group, apiContainer *container.APIContainer) {
	v1todo := g.Group("/v1")

	openapi.RegisterHandlers(v1todo, apiContainer.Todo)
}
