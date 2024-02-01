# Notenverwaltung - REST API Backend for School Grade Management
## Overview
This repository contains a REST API backend written in Golang for managing school grades. The backend utilizes the Go Echo web framework, Gorm for database connection, and Swaggo for API documentation. Authentication is done using Basic Auth.

## Features
CRUD operations for managing school grades
Authentication via Basic Auth
Automatically generated API documentation with Swaggo

## Installation
To start the server, download the latest version from the release page for your operating system. Then, run the application, and the server will be available locally at localhost:1323/api/v1.

## Example for Linux:
```bash
./Notenverwaltung
```
Example for Windows:
```powershell
.\Notenverwaltung.exe
```

## API Documentation
The available endpoints listed below are an excerpt, and the complete list can be found in the Swagger documentation at http://localhost:1323/swagger/index.html once the server is started.

### API Endpoints
- GET /api/v1/users: Retrieve all users
- GET /api/v1/students: Retrieve all students
- GET /api/v1/schools: Retrieve all schools
- GET /api/v1/classes: Retrieve all classes
- GET /api/v1/grades: Retrieve all grades
- GET /api/v1/subjects: Retrieve all subjects
- GET /api/v1/scores: Retrieve all scores
- GET /api/v1/suce: Retrieve all SUCE (replace with your specific endpoint)

For each endpoint, CRUD functionality is available.

## Authentication
The backend uses Basic Auth for authentication. Make sure to provide the appropriate credentials when calling the API.

```bash
curl -u Username:Password http://localhost:1323/auth
```
