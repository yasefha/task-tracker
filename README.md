# Task Tracker CLI
A simple command-line tool to manage your daily tasks efficiently.
Manage tasks directly from the terminal:
add, list, update, delete, and reset tasks.

## Features
- Add a new task
- List tasks (all or by status: todo | in-progress | done)
- Update task status or description
- Delete a task
- Reset all tasks

## Installation (Windows)
Download the latest Windows binary from GitHub and place it in a folder in your PATH:

### Download task.exe to user bin folder
```powershell
Invoke-WebRequest -Uri "https://github.com/yasefhatam/task-tracker/releases/latest/download/task.exe" -OutFile "$env:USERPROFILE\bin\task.exe"
```
### Add folder to PATH (temporary for current session)
```powershell
$env:PATH += ";$env:USERPROFILE\bin"
```
### Test
```powershell
task help
```

## Usage
### Add a task
```powershell
task add "Buy milk"
```
### List tasks
```powershell
task list
task list todo
```
### Update status
```powershell
task update-status 1 done
```
### Update description
```powershell
task update-desc 2 "Fix bug"
```
### Delete a task
```powershell
task delete 3
```
### Reset all tasks
```powershell
task reset --confirm
```

## Commands Overview
```
add <description>          Add a new task
list [status]              List tasks (all if no status)
update-status <id> <status> Update task status
update-desc <id> <description> Update task description
delete <id>                Delete a task by ID
reset [--confirm]          Delete all tasks permanently
```
## License
MIT License © Yasef Hatam
