package delivery

import (
	// "github.com/labstack/echo"
	"task-manager/internal/infrastructure/inmemory"
	"task-manager/internal/usecase"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	router := echo.New()

	repo := inmemory.NewTaskRepo()
	uc := usecase.NewTaskUseCase(repo)
	NewTaskHandler(router, uc)
	return router
}
