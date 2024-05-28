# Go backend clean architecture based on SQL database

## Introduction

This project is a modified version of [go-backend-clean-architecture](https://github.com/amitshekhariitbhu/go-backend-clean-architecture), changing the database layer **from mongoDB to SQL database**, plus adding some gin framework integration test examples and comments in code.

To understand clean architecture, please go to original author's blog: [Go Backend Clean Architecture](https://amitshekhar.me/blog/go-backend-clean-architecture)

## Tech Used
• Docker (Environment setup)  
• Gin framework (Http router)  
• JWT authentication  
• Mock (Unit test)  
• Postgres (Database)  
• Gorm (golang orm)  

## Installation
### Requirements: 
Install docker and docker-compose (latest version recommended).  
vscode and devcontainer extension (optional but recommended, or you want to manually run docker compose)  
### How to run in vscode
Ctrl+Shift+P to open command menu. Search and Choose **Rebuild and Reopen in container**  
After reopening in container, run test to check if everything is OK.  
```bash
go test go-backend/...
```

## Tips
You can run mock.sh in go-backend/domain to generate new mock file  
To start server  
```bash
cd ./go-backend
go run app/main.go
```

## Give a :star: if it's helpful to you



