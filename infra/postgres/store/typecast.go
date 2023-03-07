package postgres

import (
	"database/sql"
	"fmt"
	"test_app2/models"
	"time"
)

func toModelPriority(p Priority) (models.Priority, error) {
	switch p {
	case PriorityNone:
		return models.PriorityNone, nil
	case PriorityLow:
		return models.PriorityLow, nil
	case PriorityMedium:
		return models.PriorityMedium, nil
	case PriorityHigh:
		return models.PriorityHigh, nil
	}
	return models.Priority(-1), fmt.Errorf("invalid: %s", p)
}

func toSQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func toDBPriority(p models.Priority) Priority {
	switch p {
	case models.PriorityNone:
		return PriorityNone
	case models.PriorityLow:
		return PriorityLow
	case models.PriorityMedium:
		return PriorityMedium
	case models.PriorityHigh:
		return PriorityHigh
	}
	return "invalid"
}
