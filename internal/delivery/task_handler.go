package delivery

import (
	"net/http"
	"strconv"
	"task-manager/internal/domain"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	repo domain.TaskUseCase
}

func NewTaskHandler(e *echo.Echo, us domain.TaskUseCase) {
	handler := &TaskHandler{repo: us}

	t := e.Group("/tasks")
	t.POST("", handler.Create)
  t.GET("", handler.GetAll)
	t.GET("/:id", handler.GetByID)
	t.PUT("/:id", handler.Update)
	t.DELETE("/:id", handler.Delete)
}

func (h *TaskHandler) Create(c echo.Context) error {
	var task domain.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	if err := h.repo.Create(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create task"})
	}

	return c.JSON(http.StatusCreated, task)
}
func (h *TaskHandler) GetAll(c echo.Context) error {
	tasks, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid task ID"})
	}
	task, err := h.repo.GetByID(int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "task not found"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid task ID"})
	}
	var task domain.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	task.ID = int(id)
	if err := h.repo.Update(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to update task"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid task ID"})
	}
	if err := h.repo.Delete(int(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}
