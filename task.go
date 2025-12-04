package main

import "time"

type Status int

const (
	StatusPending Status = iota
	StatusInProgress
	StatusCompleted
)

type Task struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
