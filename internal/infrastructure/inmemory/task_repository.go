package inmemory

import (
	"errors"
	"sync"
	"task-manager/internal/domain"
)

type TaskRepo struct {
	tasks  map[int]domain.Task
	mu     sync.RWMutex
	nextID int
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		tasks:  make(map[int]domain.Task),
		nextID: 1,
	}
}

func (r *TaskRepo) Create(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	task.ID = r.nextID
	r.tasks[task.ID] = *task
	r.nextID++

	return nil
}

func (r *TaskRepo) GetAll() ([]domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tasks := make([]domain.Task, 0, len(r.tasks))

	for _, t := range r.tasks {
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TaskRepo) GetByID(id int) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

func (r *TaskRepo) Update(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[task.ID]; !ok {
		return errors.New("task not found")
	}
	r.tasks[task.ID] = *task
	return nil
}

func (r *TaskRepo) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
