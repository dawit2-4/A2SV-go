# Task Manager API Documentation

This document describes the RESTful API for the Task Manager application. All endpoints are prefixed with `/tasks`.

## Base URL
`http://localhost:8080`

## Endpoints

### 1. Create a Task
- **Method**: POST
- **Path**: `/tasks`
- **Description**: Creates a new task.
- **Request Body**:
  ```json
  {
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "YYYY-MM-DDTHH:MM:SSZ",
    "status": "pending|completed|not-done"
  }
  ```
- **Response**:
  - **201 Created**: Task created successfully.
    ```json
    {
      "id": "1",
      "title": "Sample Task",
      "description": "Test task",
      "due_date": "2025-07-18T12:00:00Z",
      "status": "pending"
    }
    ```
  - **400 Bad Request**: Invalid request body or validation error (e.g., empty title, invalid status).
    ```json
    {"error": "task title cannot be empty"}
    ```

### 2. Get a Task by ID
- **Method**: GET
- **Path**: `/tasks/:id`
- **Description**: Retrieves a task by its ID.
- **Response**:
  - **200 OK**: Task found.
    ```json
    {
      "id": "1",
      "title": "Sample Task",
      "description": "Test task",
      "due_date": "2025-07-18T12:00:00Z",
      "status": "pending"
    }
    ```
  - **404 Not Found**: Task not found.
    ```json
    {"error": "task with ID 1 not found"}
    ```

### 3. Get All Tasks
- **Method**: GET
- **Path**: `/tasks`
- **Description**: Retrieves all tasks.
- **Response**:
  - **200 OK**: List of tasks.
    ```json
    [
      {
        "id": "1",
        "title": "Sample Task",
        "description": "Test task",
        "due_date": "2025-07-18T12:00:00Z",
        "status": "pending"
      }
    ]
    ```
  - **500 Internal Server Error**: Unexpected error.
    ```json
    {"error": "internal server error"}
    ```

### 4. Update a Task
- **Method**: PUT
- **Path**: `/tasks/:id`
- **Description**: Updates an existing task.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "YYYY-MM-DDTHH:MM:SSZ",
    "status": "pending|completed|not-done"
  }
  ```
- **Response**:
  - **200 OK**: Task updated successfully.
    ```json
    {
      "id": "1",
      "title": "Updated Task",
      "description": "Updated description",
      "due_date": "2025-07-18T12:00:00Z",
      "status": "completed"
    }
    ```
  - **400 Bad Request**: Invalid request body or validation error.
    ```json
    {"error": "task title cannot be empty"}
    ```
  - **404 Not Found**: Task not found.
    ```json
    {"error": "task with ID 1 not found"}
    ```

### 5. Delete a Task
- **Method**: DELETE
- **Path**: `/tasks/:id`
- **Description**: Deletes a task by its ID.
- **Response**:
  - **204 No Content**: Task deleted successfully.
  - **404 Not Found**: Task not found.
    ```json
    {"error": "task with ID 1 not found"}
    ```

## Error Handling
- All errors are returned in JSON format with an `error` field describing the issue.
- Common HTTP status codes:
  - `200 OK`: Successful request.
  - `201 Created`: Resource created.
  - `204 No Content`: Successful deletion.
  - `400 Bad Request`: Invalid input or validation error.
  - `404 Not Found`: Resource not found.
  - `500 Internal Server Error`: Unexpected server error.

## Example Usage
Create a task:
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Sample Task","description":"Test task","due_date":"2025-07-18T12:00:00Z","status":"pending"}'
```

Get a task:
```bash
curl http://localhost:8080/tasks/1
```