# Task Manager API

## Overview

The Task Manager API is a RESTful web service built with Go, using the Gin framework for routing and MongoDB for data persistence. It provides functionality for managing tasks and user authentication, with support for role-based access control (Admin and User roles). The API supports CRUD operations for tasks and user registration/login with JWT-based authentication.

## Features

- **Task Management**: Create, read, update, and delete tasks with title, description, due date, and status.
- **User Authentication**: Register and login users, with password hashing and JWT token generation.
- **Role-Based Access Control**: Admin users can delete tasks, while regular users can perform other task operations.
- **MongoDB Integration**: Stores tasks and user data with efficient indexing.
- **Secure Authentication**: Uses JWT for protecting routes and bcrypt for password hashing.

## Project Structure

The application is organized into packages for separation of concerns:

- **`main.go`**: Initializes the server, connects to MongoDB, and sets up routes.
- **`routers/router.go`**: Configures HTTP routes for tasks and authentication.
- **`models/task.go`**: Defines the `Task` struct and task status with validation.
- **`models/setup.go`**: Handles MongoDB connection and index creation.
- **`models/user.go`**: Defines the `User` struct and user roles with validation.
- **`middlewares/middleware.go`**: Implements JWT authentication and admin-only middleware.
- **`data/user_service.go`**: Manages user-related database operations (register, login, retrieve).
- **`data/task_service.go`**: Manages task-related database operations (CRUD).
- **`controllers/task_controller.go`**: Handles HTTP requests for task and user operations.

## Setup Instructions

### Prerequisites

- **Go**: Version 1.16 or higher.
- **MongoDB**: Running locally or accessible via a connection string.
- **Environment Variables**:
  - `MONGODB_URI`: MongoDB connection string (default: `mongodb://localhost:27017`).
  - `JWT_SECRET`: Secret key for JWT signing (e.g., a random, secure string).

### Installation

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd task_manager
   ```

2. **Install Dependencies**:
   ```bash
   go get github.com/gin-gonic/gin
   go get go.mongodb.org/mongo-driver/mongo
   go get github.com/dgrijalva/jwt-go
   go get golang.org/x/crypto/bcrypt
   ```

3. **Set Environment Variables**:
   ```bash
   export MONGODB_URI="mongodb://localhost:27017"
   export JWT_SECRET="your-secure-jwt-secret"
   ```

4. **Run the Application**:
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080`.

## API Endpoints

### Authentication Routes

- **POST /register**
  - **Description**: Register a new user.
  - **Request Body**:
    ```json
    {
      "username": "string",
      "password": "string",
      "role": "Admin|User"
    }
    ```
  - **Response**:
    - `201 Created`: `{ "message": "user created successfully", "user": { "id": "string", "username": "string", "role": "string" } }`
    - `400 Bad Request`: If username is taken or input is invalid.
    - `500 Internal Server Error`: On server errors.
  - **Example**:
    ```bash
    curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username":"john","password":"secure123","role":"User"}'
    ```

- **POST /login**
  - **Description**: Authenticate a user and return a JWT token.
  - **Request Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Response**:
    - `200 OK`: `{ "token": "string" }`
    - `401 Unauthorized`: If credentials are invalid.
    - `500 Internal Server Error`: On server errors.
  - **Example**:
    ```bash
    curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username":"john","password":"secure123"}'
    ```

### Task Routes (Protected)

All task routes require a valid JWT token in the `Authorization` header (`Bearer <token>`).

- **POST /tasks**
  - **Description**: Create a new task.
  - **Request Body**:
    ```json
    {
      "title": "string",
      "description": "string",
      "due_date": "2025-12-31T23:59:59Z",
      "status": "pending|completed|not-done"
    }
    ```
  - **Response**:
    - `201 Created`: Returns the created task.
    - `400 Bad Request`: If input is invalid.
  - **Example**:
    ```bash
    curl -X POST http://localhost:8080/tasks -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title":"Finish report","description":"Complete quarterly report","due_date":"2025-12-31T23:59:59Z","status":"pending"}'
    ```

- **GET /tasks**
  - **Description**: Retrieve all tasks.
  - **Response**:
    - `200 OK`: Returns a list of tasks.
    - `500 Internal Server Error`: On server errors.
  - **Example**:
    ```bash
    curl -X GET http://localhost:8080/tasks -H "Authorization: Bearer <token>"
    ```

- **GET /tasks/:id**
  - **Description**: Retrieve a task by ID.
  - **Response**:
    - `200 OK`: Returns the task.
    - `400 Bad Request`: If ID is invalid or task not found.
  - **Example**:
    ```bash
    curl -X GET http://localhost:8080/tasks/507f1f77bcf86cd799439011 -H "Authorization: Bearer <token>"
    ```

- **PUT /tasks/:id**
  - **Description**: Update a task by ID.
  - **Request Body**: Same as POST /tasks.
  - **Response**:
    - `200 OK`: Returns the updated task.
    - `400 Bad Request`: If ID or input is invalid.
  - **Example**:
    ```bash
    curl -X PUT http://localhost:8080/tasks/507f1f77bcf86cd799439011 -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title":"Updated report","description":"Revised report","due_date":"2025-12-31T23:59:59Z","status":"completed"}'
    ```

- **DELETE /tasks/:id**
  - **Description**: Delete a task by ID (admin only).
  - **Response**:
    - `200 OK`: `{ "message": "task deleted successfully" }`
    - `400 Bad Request`: If ID is invalid or task not found.
    - `403 Forbidden`: If user is not an admin.
  - **Example**:
    ```bash
    curl -X DELETE http://localhost:8080/tasks/507f1f77bcf86cd799439011 -H "Authorization: Bearer <token>"
    ```

## Data Models

### Task
```json
{
  "id": "string", // MongoDB ObjectID
  "title": "string", // Required, max 100 characters
  "description": "string", // Optional, max 1000 characters
  "due_date": "string", // ISO 8601 format, must be future date
  "status": "pending|completed|not-done" // Required
}
```

### User
```json
{
  "id": "string", // MongoDB ObjectID
  "username": "string", // Required, max 50 characters
  "password": "string", // Required, min 8 characters (hashed in database)
  "role": "Admin|User" // Required
}
```

## Security

- **JWT Authentication**: Task routes require a valid JWT token in the `Authorization` header.
- **Admin Access**: Only users with the `Admin` role can delete tasks.
- **Password Hashing**: Passwords are hashed using bcrypt before storage.
- **MongoDB Indexes**: Unique index on `username` ensures no duplicates.

## Error Responses

Common error responses include:
- `400 Bad Request`: Invalid input or ID format.
- `401 Unauthorized`: Missing or invalid JWT token, or invalid credentials.
- `403 Forbidden`: Non-admin user attempting admin-only actions.
- `500 Internal Server Error`: Unexpected server errors.

Example:
```json
{
  "error": "invalid input data: title cannot be empty"
}
```

## Environment Variables

- `MONGODB_URI`: MongoDB connection string (e.g., `mongodb://localhost:27017`).
- `JWT_SECRET`: Secret key for JWT signing (e.g., a 32-character random string).

## Running Tests

To add tests, use Go's `testing` package. Example test setup:
```bash
go test ./...
```
Currently, no tests are included. Consider adding unit tests for `data` and `controllers` packages using `github.com/stretchr/testify`.

## Future Improvements

- **Pagination**: Add pagination to `GET /tasks` for large datasets.
- **CORS**: Add CORS middleware for frontend integration.
- **Rate Limiting**: Implement rate limiting to prevent abuse.
- **Logging**: Use a structured logger (e.g., `github.com/sirupsen/logrus`).
- **HTTPS**: Deploy with TLS for secure communication.

## Example Usage

1. **Register a User**:
   ```bash
   curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username":"admin","password":"securepassword123","role":"Admin"}'
   ```

2. **Login to Get JWT Token**:
   ```bash
   curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username":"admin","password":"securepassword123"}'
   ```

3. **Create a Task**:
   ```bash
   curl -X POST http://localhost:8080/tasks -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title":"Write code","description":"Write Go API","due_date":"2025-12-31T23:59:59Z","status":"pending"}'
   ```

4. **Delete a Task (Admin Only)**:
   ```bash
   curl -X DELETE http://localhost:8080/tasks/507f1f77bcf86cd799439011 -H "Authorization: Bearer <token>"
   ```

## License

This project is unlicensed. Use and modify as needed.