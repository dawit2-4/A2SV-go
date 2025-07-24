# Task Manager API (Clean Architecture)

## Overview

The Task Manager API is a RESTful web service built with Go, using Clean Architecture principles to ensure maintainability, testability, and scalability. It provides task management (CRUD operations) and user authentication (register/login) with JWT-based authentication and role-based access control (Admin/User). The application uses MongoDB for data persistence and the Gin framework for HTTP routing.

## Architecture

The codebase follows Clean Architecture, with layers separated by responsibility:

- **Domain**: Core business entities (`Task`, `User`) and repository interfaces, independent of frameworks.
- **Usecases**: Business logic for task and user operations, orchestrating interactions with repositories.
- **Repositories**: Implements data access with MongoDB, adhering to Domain interfaces.
- **Infrastructure**: External services (JWT, password hashing, middleware).
- **Delivery**: HTTP handlers and routers, interacting with use cases.

### Folder Structure

```
task-manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleware.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/
│   ├── task_repository.go
│   └── user_repository.go
├── Usecases/
│   ├── task_usecases.go
│   └── user_usecases.go
├── task_manager_test.go
├── .env
├── README.md
```

## Features

- **Task Management**: Create, read, update, and delete tasks with title, description, due date, and status.
- **User Authentication**: Register and login users with JWT tokens and bcrypt password hashing.
- **Role-Based Access**: Admins can delete tasks; all users can perform other operations.
- **Clean Architecture**: Layered design with clear separation of concerns and dependency inversion.
- **MongoDB Integration**: Efficient data storage with indexing.
- **Unit Tests**: Tests for use cases and controllers using mocks.

## Setup Instructions

### Prerequisites

- **Go**: Version 1.16 or higher.
- **MongoDB**: Running locally or accessible via a connection string.
- **Environment Variables**:
  - `MONGODB_URI`: MongoDB connection string (default: `mongodb://localhost:27017`).
  - `JWT_SECRET`: Secret key for JWT signing.
  - `DB_NAME`: MongoDB database name (default: `tasks`).
  - `TASKS_COLLECTION`: MongoDB collection for tasks (default: `tasks`).
  - `USERS_COLLECTION`: MongoDB collection for users (default: `users`).

### Installation

1. **Clone the Repository**:

   ```bash
   git clone <repository-url>
   cd task-manager
   ```

2. **Install Dependencies**:

   ```bash
   go get github.com/gin-gonic/gin
   go get go.mongodb.org/mongo-driver/mongo
   go get github.com/dgrijalva/jwt-go
   go get golang.org/x/crypto/bcrypt
   go get github.com/stretchr/testify
   go get github.com/joho/godotenv
   ```

3. **Create a .env File**:
   Create a `.env` file in the project root with the following content:

   ```env
   MONGODB_URI=mongodb://localhost:27017
   JWT_SECRET=your-secure-jwt-secret
   DB_NAME=tasks
   TASKS_COLLECTION=tasks
   USERS_COLLECTION=users
   ```

   Ensure the `.env` file is added to `.gitignore` to avoid exposing sensitive data.

4. **Set Environment Variables (Optional)**:
   Alternatively, set the variables directly:

   ```bash
   export MONGODB_URI="mongodb://localhost:27017"
   export JWT_SECRET="your-secure-jwt-secret"
   export DB_NAME="tasks"
   export TASKS_COLLECTION="tasks"
   export USERS_COLLECTION="users"
   ```

5. **Run the Application**:
   ```bash
   go run Delivery/main.go
   ```
   The server runs on `http://localhost:8080`.

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
    - `400 Bad Request`: Invalid input or username taken.
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
    - `401 Unauthorized`: Invalid credentials.
  - **Example**:
    ```bash
    curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username":"john","password":"secure123"}'
    ```

### Task Routes (Protected)

Require `Authorization: Bearer <token>` header.

- **POST /tasks**

  - **Description**: Create a task.
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
    - `201 Created`: Task object.
    - `400 Bad Request`: Invalid input.
  - **Example**:
    ```bash
    curl -X POST http://localhost:8080/tasks -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title":"Finish report","description":"Complete quarterly report","due_date":"2025-12-31T23:59:59Z","status":"pending"}'
    ```

- **GET /tasks**

  - **Description**: Retrieve all tasks.
  - **Response**:
    - `200 OK`: List of tasks.
    - `500 Internal Server Error`: Server error.
  - **Example**:
    ```bash
    curl -X GET http://localhost:8080/tasks -H "Authorization: Bearer <token>"
    ```

- **GET /tasks/:id**

  - **Description**: Retrieve a task by ID.
  - **Response**:
    - `200 OK`: Task object.
    - `400 Bad Request`: Invalid ID or not found.
  - **Example**:
    ```bash
    curl -X GET http://localhost:8080/tasks/507f1f77bcf86cd799439011 -H "Authorization: Bearer <token>"
    ```

- **PUT /tasks/:id**

  - **Description**: Update a task.
  - **Request Body**: Same as POST /tasks.
  - **Response**:
    - `200 OK`: Updated task.
    - `400 Bad Request`: Invalid input or ID.
  - **Example**:
    ```bash
    curl -X PUT http://localhost:8080/tasks/507f1f77bcf86cd799439011 -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title":"Updated report","description":"Revised report","due_date":"2025-12-31T23:59:59Z","status":"completed"}'
    ```

- **DELETE /tasks/:id**
  - **Description**: Delete a task (admin only).
  - **Response**:
    - `200 OK`: `{ "message": "task deleted successfully" }`
    - `400 Bad Request`: Invalid ID or not found.
    - `403 Forbidden`: Non-admin user.
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
  "due_date": "string", // ISO 8601, future date
  "status": "pending|completed|not-done" // Required
}
```

### User

```json
{
  "id": "string", // MongoDB ObjectID
  "username": "string", // Required, max 50 characters
  "password": "string", // Required, min 8 characters (hashed)
  "role": "Admin|User" // Required
}
```

## Running Tests

Unit tests are provided in `task_manager_test.go` using `testify/mock`.

1. **Install Test Dependencies**:

   ```bash
   go get github.com/stretchr/testify
   ```

2. **Run Tests**:
   ```bash
   go test -v ./...
   ```

Tests cover:

- `TaskUsecase`: CreateTask (valid/invalid input).
- `UserUsecase`: RegisterUser, LogIn.
- `TaskController`: CreateTask.

## Design Decisions

- **Clean Architecture**: Layers are isolated, with dependencies flowing inward (Delivery -> Usecases -> Domain).
- **Dependency Inversion**: Repository interfaces in `Domain`, implemented in `Repositories`. Infrastructure services (JWT, password) are abstracted via interfaces.
- **Domain Independence**: `Domain` package contains pure Go structs and interfaces, free of external dependencies except `mongo-driver` for ObjectID.
- **Configuration**: Environment variables (`MONGODB_URI`, `JWT_SECRET`, `DB_NAME`, `TASKS_COLLECTION`, `USERS_COLLECTION`) are loaded from a `.env` file or environment, with defaults for flexibility.
- **Simplified Middleware**: Moved to `Infrastructure`, with `JWTService` and `PasswordService` as abstractions.
- **Backward Compatibility**: API endpoints and functionality match the original implementation.
- **MongoDB Indexes**: Ensured for performance and uniqueness.

## Future Improvements

- **Pagination**: Add to `GetAllTasks` for large datasets.
- **CORS**: Enable for frontend integration.
- **Rate Limiting**: Prevent API abuse.
- **Structured Logging**: Use `logrus` for better logging.
- **HTTPS**: Deploy with TLS.

## License

Unlicensed. Use and modify as needed.
