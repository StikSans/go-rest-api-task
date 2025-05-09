package usecase

import "task-manager/internal/domain"

type taskUseCase struct {
	repo domain.TaskRepository
}

func NewTaskUseCase(repo domain.TaskRepository) domain.TaskUseCase {
	return &taskUseCase{repo: repo}
}


func (u *taskUseCase) Create(task *domain.Task) error {
	return u.repo.Create(task)
}
func (u *taskUseCase) GetAll() ([]domain.Task, error) {
	return u.repo.GetAll()
}
func (u *taskUseCase) GetByID(id int) (*domain.Task, error) {
	return u.repo.GetByID(id)
}
func (u *taskUseCase) Update(task *domain.Task) error {
	return u.repo.Update(task)
}
func (u *taskUseCase) Delete(id int) error {
	return u.repo.Delete(id)
}