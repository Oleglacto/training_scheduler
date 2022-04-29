package models

import "time"

type Training struct {
	ID          string
	Name        string
	City        string
	Description string
	Place       string
	StartAt     time.Time
	EndAt       time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
