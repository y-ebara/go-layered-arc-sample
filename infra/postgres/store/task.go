package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"test_app2/models"

	"github.com/google/uuid"
)

type taskRepo struct {
	q *Queries
}

func NewTaskRepository(db *sql.DB) *taskRepo {
	return &taskRepo{
		q: New(db),
	}
}

func (t *taskRepo) Create(ctx context.Context, description string, priority models.Priority, period models.TaskPeriod) (models.Task, error) {
	id, err := t.q.CreateTask(ctx, CreateTaskParams{
		Description: description,
		Priority:    toDBPriority(priority),
		StartDate:   toSQLNullTime(period.Start),
		DueDate:     toSQLNullTime(period.Limit),
	})
	if err != nil {
		return models.Task{}, fmt.Errorf("err creating task: %w", err)
	}

	return models.Task{
		ID:          id.String(),
		Description: description,
		Priority:    priority,
		Period:      period,
	}, nil
}

func (t *taskRepo) Get(ctx context.Context, id string) (models.Task, error) {
	uc, err := uuid.Parse(id)
	if err != nil {
		return models.Task{}, fmt.Errorf("err parsing id: %w", err)
	}
	task, err := t.q.GetTask(ctx, uc)
	if err != nil {
		return models.Task{}, fmt.Errorf("err getting task: %w", err)
	}
	p, err := toModelPriority(task.Priority)
	if err != nil {
		return models.Task{}, fmt.Errorf("err cast priority model: %w", err)
	}
	return models.Task{
		ID:          task.ID.String(),
		Description: task.Description,
		Priority:    p,
		Period: models.TaskPeriod{
			Start: task.StartDate.Time,
			Limit: task.DueDate.Time,
		},
	}, nil
}



func (t *taskRepo) Update(ctx context.Context, id string, description string, priority  models.Priority, period models.TaskPeriod, completed bool) error {
	uc, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("err parsing id: %w", err)
	}
	if err := t.q.UpdateTask(ctx, UpdateTaskParams{
		ID:          uc,
		Description: description,
		Priority:    toDBPriority(priority),
		StartDate:   toSQLNullTime(period.Start),
		DueDate:     toSQLNullTime(period.Limit),
		Done:        completed,
	}); err != nil {
		return fmt.Errorf("err updating task: %w", err)
	}
	return nil
}

func (t *taskRepo) GetDownTask(ctx context.Context, id string) (models.DoneTask, error) {
	return models.DoneTask{}, nil
}

