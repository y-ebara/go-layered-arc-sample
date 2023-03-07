package inmemory

import (
	"context"
	"errors"
	"test_app2/models"
	"test_app2/repository"

	"github.com/google/uuid"
)

type inMemory struct {
	store map[string]models.Task
}

func NewInMemory() repository.TaskRepository {
	return &inMemory{store: make(map[string]models.Task)}
}

func (i *inMemory) Create(ctx context.Context, description string, priority models.Priority, dates models.TaskPeriod) (models.Task, error) {
	task := models.Task{
		ID:          uuid.New().String(),
		Description: description,
		Priority:    priority,
		Period:      dates,
		Completed:   false,
	}
	i.store[task.ID] = task
	return task, nil
}

func (i *inMemory) Get(ctx context.Context, id string) (models.Task, error) {
	task, ok := i.store[id]
	if !ok {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (i *inMemory) Update(ctx context.Context, id string, description string, priority models.Priority, dates models.TaskPeriod, completed bool) error {
	task, ok := i.store[id]
	if !ok {
		return errors.New("task not found")
	}
	task.Description = description
	task.Priority = priority
	task.Period = dates
	task.Completed = completed
	i.store[id] = task
	return nil
}

func (t *inMemory) GetDownTask(ctx context.Context, id string) (models.DoneTask, error) {
	return models.DoneTask{}, nil
}

