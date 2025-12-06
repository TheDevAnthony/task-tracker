# 📝 task-tracker

A simple, lightweight CLI tool for managing tasks using JSON storage.

## 🚀 Overview

`taskcli` lets you add, update, delete, and track tasks directly from the terminal.  
All data is stored locally in a readable `tasks.json` file — no setup or dependencies needed.

## ✨ Features

- Add, update, and delete tasks  
- Mark tasks as **todo**, **in-progress**, or **done**  
- List all tasks or filter by status  
- JSON file persistence  
- Pure Go, no external dependencies  

## 🎯 Usage

#### Add, update and delete tasks
```bash
task-cli add "Buy groceries"
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1
```
#### Change task status
```bash
task-cli mark-in-progress 1
task-cli mark-done 1
```
#### List tasks
```bash
task-cli list
task-cli list done
task-cli list todo
task-cli list in-progress
```

## 📦 Installation
For the installation to work, make sure you have Go installed on your machine, then just run in the cmd:
```bash
go install github.com/TheDevAnthony/task-tracker/v2/cmd/taskcli@latest
```
<hr>
This project is actually my solution for the <a href="https://roadmap.sh/projects/task-tracker">Backend Roadmap - Task Tracker Project</a>
