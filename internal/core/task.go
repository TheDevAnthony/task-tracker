package core

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in-progress"
	StatusCompleted  Status = "completed"
)

type Task struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var tasks []Task
var dataFile = filepath.Join(".", "tasks.json")

func LoadTasks() error {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			tasks = []Task{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &tasks)
}

func SaveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}
