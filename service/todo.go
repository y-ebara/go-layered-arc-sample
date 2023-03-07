package service

import (
	"context"
	"test_app2/models"
	"test_app2/repository"
)


type TodoService interface {
	CreateNewTask(ctx context.Context, description string, priority models.Priority, period models.TaskPeriod) (models.Task, error)
	FindTask(ctx context.Context, id string) (models.Task, error)
	UpdateTask(ctx context.Context, id string, description string, priority models.Priority, period models.TaskPeriod, completed bool) error
	GetDoneTask(ctx context.Context, id string) (models.DoneTask, error)
}

type todoService struct {
	Repo repository.TaskRepository
}

func NewTodoService(repo repository.TaskRepository) TodoService {
	return &todoService{Repo: repo}
}

func (t *todoService) CreateNewTask(ctx context.Context, description string, priority models.Priority, period models.TaskPeriod) (models.Task, error) {
	err := period.Validate()
	if err != nil {
		return models.Task{}, err
	}
	err = priority.Validate()
	if err != nil {
		return models.Task{}, err
	}
	task, err := t.Repo.Create(ctx, description, priority, period)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *todoService) FindTask(ctx context.Context, id string) (models.Task, error) {
	task, err := t.Repo.Get(ctx, id)
	if err != nil {
		return models.Task{}, err
	}
	err = task.Validate()
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *todoService) UpdateTask(ctx context.Context, id string, description string, priority models.Priority, period models.TaskPeriod, completed bool) error {
	err := period.Validate()
	if err != nil {
		return err
	}
	err = priority.Validate()
	if err != nil {
		return err
	}
	err = t.Repo.Update(ctx, id, description, priority, period, completed)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoService) GetDoneTask(ctx context.Context, id string) (models.DoneTask, error) {
	task, err := t.Repo.GetDownTask(ctx, id)
	if err != nil {
		return models.DoneTask{}, err
	}
	return task, nil
}
