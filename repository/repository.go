package repository

import (
	"context"
	"test_app2/models"
)

type TaskRepository interface {
	Create(ctx context.Context, description string, priority models.Priority, dates models.TaskPeriod) (models.Task, error)
	Get(ctx context.Context, id string) (models.Task, error)
	Update(ctx context.Context, id string, description string, priority models.Priority, dates models.TaskPeriod, completed bool) error
	GetDownTask(ctx context.Context, id string) (models.DoneTask, error)
}