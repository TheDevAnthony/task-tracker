# 📝 task-tracker

A simple, lightweight CLI tool for managing tasks using JSON storage.

## 🚀 Overview

`taskcli` lets you add, update, delete, and track tasks directly from the terminal.  
All data is stored locally in a readable `tasks.json` file — no setup or dependencies needed.

## ✨ Features

- Add, update, and delete tasks  
- Mark tasks as **in-progress** or **completed**  
- List all tasks or filter by status  
- JSON file persistence  
- Pure Go, no external dependencies  

## 🎯 Usage

#### Add, update and delete tasks
```bash
taskcli add "Buy groceries"
taskcli update 1 "Buy groceries and cook dinner"
taskcli delete 1
```
#### Change task status
```bash
taskcli mark-in-progress 1
taskcli mark-completed 1
```
#### List tasks
```bash
taskcli list
taskcli list completed
taskcli list pending
taskcli list in-progress
```

## 📦 Installation
For the installation to work, make sure you have Go installed on your machine, then just run in the cmd:
```bash
go install github.com/TheDevAnthony/task-tracker/v2/cmd/taskcli@latest
```
<hr>
This project is actually my solution for the <a href="https://roadmap.sh/projects/task-tracker">Backend Roadmap - Task Tracker Project</a>
