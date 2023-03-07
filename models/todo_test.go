package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_PriorityValidate(t *testing.T) {
	tests := []struct {
		name string
		p    Priority
		want bool
	}{
		{"none", PriorityNone, false},
		{"low", PriorityLow, false},
		{"medium", PriorityMedium, false},
		{"high", PriorityHigh, false},
		{"invalid", Priority(4), true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.p.Validate()
			assert.Equal(t, tc.want, err != nil)
		})
	}
}

func Test_TaskPeriodValidate(t *testing.T) {
	tests := []struct {
		name    string
		tp      TaskPeriod
		wantErr bool
	}{
		{"invalid", TaskPeriod{time.Now().Add(time.Hour), time.Now()}, true},
		{"valid", TaskPeriod{time.Now(), time.Now().Add(time.Hour)}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.tp.Validate()
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func Test_TaskValidate(t *testing.T) {
	tests := []struct {
		name    string
		t       Task
		wantErr bool
	}{
		{"valid", Task{Period: TaskPeriod{time.Now(), time.Now().Add(time.Hour)}, Priority: PriorityLow}, false},
		{"invalid when period is invalid", Task{Period: TaskPeriod{time.Now().Add(time.Hour), time.Now()}, Priority: PriorityLow}, true},
		{"invalid when period is invalid", Task{Period: TaskPeriod{time.Now(), time.Now().Add(time.Hour)}, Priority: Priority(10)}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.t.Validate()
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
