package domain

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	CreatedAt   string `json:"createdAt"`
}
type TaskRepository interface {
	Create(task *Task) error
	GetAll() ([]Task, error)
	GetByID(id int) (*Task, error)
	Update(task *Task) error
	Delete(id int) error
}
type TaskUseCase interface {
	Create(task *Task) error
	GetAll() ([]Task, error)
	GetByID(id int) (*Task, error)
	Update(task *Task) error
	Delete(id int) error
}
