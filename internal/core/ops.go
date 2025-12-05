package core

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func AddTask(desc string) {
	if strings.TrimSpace(desc) == "" {
		log.Fatal("description cannot be empty")
	}

	task := Task{
		ID:          len(tasks) + 1,
		Description: desc,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	fmt.Printf("Added task: %d - %s\n", task.ID, task.Description)

	if err := SaveTasks(); err != nil {
		log.Fatalf("failed to save tasks: %v", err)
	}
}

func UpdateTask(str_id string, desc string) {
	id, err := strconv.Atoi(str_id)
	if err != nil {
		log.Fatal("invalid task ID")
	}

	if strings.TrimSpace(desc) == "" {
		log.Fatal("description cannot be empty")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now()
			fmt.Printf("Updated task: %d - %s\n", tasks[i].ID, tasks[i].Description)
			if err := SaveTasks(); err != nil {
				log.Fatalf("failed to save tasks: %v", err)
			}
			return
		}
	}
	log.Fatal("task not found")
}

func DeleteTask(str_id string) {
	id, err := strconv.Atoi(str_id)
	if err != nil {
		log.Fatal("invalid task ID")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Deleted task: %d\n", id)
			if err := SaveTasks(); err != nil {
				log.Fatalf("failed to save tasks: %v", err)
			}
			return
		}
	}
	log.Fatal("task not found")
}

func MarkInProgress(str_id string) {
	id, err := strconv.Atoi(str_id)
	if err != nil {
		log.Fatal("invalid task ID")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = StatusInProgress
			tasks[i].UpdatedAt = time.Now()
			fmt.Printf("Marked task as in-progress: %d\n", tasks[i].ID)
			if err := SaveTasks(); err != nil {
				log.Fatalf("failed to save tasks: %v", err)
			}
			return
		}
	}
	log.Fatal("task not found")
}

func MarkCompleted(str_id string) {
	id, err := strconv.Atoi(str_id)
	if err != nil {
		log.Fatal("invalid task ID")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = StatusCompleted
			tasks[i].UpdatedAt = time.Now()
			fmt.Printf("Marked task as completed: %d\n", tasks[i].ID)
			if err := SaveTasks(); err != nil {
				log.Fatalf("failed to save tasks: %v", err)
			}
			return
		}
	}
	log.Fatal("task not found")
}

func ListTasks(status *string) {
	var st string
	if status != nil {
		st = *status
	} else {
		st = "all"
	}

	switch st {
	case "all":
		fmt.Println("Listing all tasks:")
		for _, task := range tasks {
			fmt.Printf("%d - %s [%s]\n", task.ID, task.Description, task.Status)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
		}
	case "pending":
		fmt.Println("Listing pending tasks:")
		for _, task := range tasks {
			if task.Status == StatusPending {
				fmt.Printf("%d - %s [%s]\n", task.ID, task.Description, task.Status)
			}
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
		}
	case "in-progress":
		fmt.Println("Listing in-progress tasks:")
		for _, task := range tasks {
			if task.Status == StatusInProgress {
				fmt.Printf("%d - %s [%s]\n", task.ID, task.Description, task.Status)
			}
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
		}
	case "completed":
		fmt.Println("Listing completed tasks:")
		for _, task := range tasks {
			if task.Status == StatusCompleted {
				fmt.Printf("%d - %s [%s]\n", task.ID, task.Description, task.Status)
			}
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
		}
	default:
		log.Fatal(errors.New("invalid status filter"))
	}
}

func PrintHelp() {
	fmt.Println(`task tracker cli - A simple task management CLI tool

Usage:
  taskcli add <description>               Add a new task
  taskcli update <id> <new_description>   Update an existing task
  taskcli delete <id>                     Delete a task
  taskcli mark-in-progress <id>           Mark a task as in-progress
  taskcli mark-completed <id>             Mark a task as completed
  taskcli list <status>                   List tasks, optionally filtered by status (pending, in-progress, completed)
  taskcli help                            Show this help message`)
}
