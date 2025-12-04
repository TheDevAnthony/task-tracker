package main

import "os"

var tasks []Task

func main() {
	if len(os.Args) < 2 {
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		if len(args) < 1 {
			return
		}
		AddTask(args[0])
	case "update":
		if len(args) < 2 {
			return
		}
		UpdateTask(args[0], args[1])
	case "delete":
		if len(args) < 1 {
			return
		}
		DeleteTask(args[0])
	case "mark-in-progress":
		if len(args) < 1 {
			return
		}
		MarkInProgress(args[0])
	case "mark-completed":
		if len(args) < 1 {
			return
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
