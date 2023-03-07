package models

import (
	"errors"
	"time"
)

type Priority int

const (
	PriorityNone Priority = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
)

func (p Priority) Validate() error {
	if p != PriorityNone && p != PriorityLow && p != PriorityMedium && p != PriorityHigh {
		return errors.New("invalid priority")
	}
	return nil
}

type TaskCategory string

type TaskPeriod struct {
	Start time.Time
	Limit time.Time
}

func (tp TaskPeriod) Validate() error {
	if tp.Start.After(tp.Limit) {
		return errors.New("invalid task period")
	}
	return nil
}

type Task struct {
	ID          string
	Description string
	Priority    Priority
	Period      TaskPeriod
	SubTask     []Task
	Category    []TaskCategory
	Completed   bool
}

func (t Task) Validate() error {
	if err := t.Priority.Validate(); err != nil {
		return err
	}
	if err := t.Period.Validate(); err != nil {
		return err
	}
	return nil
}

type DoneTask struct {
	ID          string
	Description string
	Category    []TaskCategory
}