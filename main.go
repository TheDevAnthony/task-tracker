package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No command provided. Use 'help' to see available commands.")
	}

	if err := LoadTasks(); err != nil {
		log.Fatalf("Error loading tasks: %v", err)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		if len(args) < 1 {
			log.Fatal("Invalid command. Usage: add <description>")
		}
		AddTask(args[0])
	case "update":
		if len(args) < 2 {
			log.Fatal("Invalid command. Usage: update <id> <new description>")
		}
		UpdateTask(args[0], args[1])
	case "delete":
		if len(args) < 1 {
			log.Fatal("Invalid command. Usage: delete <id>")
		}
		DeleteTask(args[0])
	case "mark-in-progress":
		if len(args) < 1 {
			log.Fatal("Invalid command. Usage: mark-in-progress <id>")
		}
		MarkInProgress(args[0])
	case "mark-completed":
		if len(args) < 1 {
			log.Fatal("Invalid command. Usage: mark-completed <id>")
		}
		MarkCompleted(args[0])
	case "list":
		var status *string
		if len(args) >= 1 {
			status = &args[0]
		}
		ListTasks(status)
	case "help":
		PrintHelp()
	default:
		return
	}
}
