package main

import (
	"log"
	"os"

	"github.com/TheDevAnthony/task-tracker/v2/internal/core"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No command provided. Use 'help' to see available commands.")
	}

	if err := core.LoadTasks(); err != nil {
		log.Fatalf("Error loading tasks: %v", err)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		if len(args) < 1 {
			log.Fatal("Invalid command. Usage: add <description>")
		}
		core.AddTask(args[0])
	case "update":
		if len(args) < 2 {
			log.Fatal("Invalid command. Usage: update <id> <new description>")
		}
		core.UpdateTask(args[0], args[1])
	case "delete":
		if len(args) < 1 {
			log.Fatal("Invalid command. Usage: delete <id>")
		}
		core.DeleteTask(args[0])
	case "mark-in-progress":
		if len(args) < 1 {
			log.Fatal("Invalid command. Usage: mark-in-progress <id>")
		}
		core.MarkInProgress(args[0])
	case "mark-completed":
		if len(args) < 1 {
			log.Fatal("Invalid command. Usage: mark-completed <id>")
		}
		core.MarkCompleted(args[0])
	case "list":
		var status *string
		if len(args) >= 1 {
			status = &args[0]
		}
		core.ListTasks(status)
	case "help":
		core.PrintHelp()
	default:
		return
	}
}
