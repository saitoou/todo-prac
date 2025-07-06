package routes

import (
	"todo-golang/infrastructure/container"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, apiContainer *container.APIContainer) {

	// TODO Custom Middleware
	registerAPIRoutes(e, apiContainer)
}
