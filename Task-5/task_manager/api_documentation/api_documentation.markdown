# ✅ Task Manager API Documentation

This document describes the RESTful API for the Task Manager application. All endpoints are prefixed with `/tasks`.

## 🔗 Base URL

```
http://localhost:8080
```

## 📌 Endpoints

### 1. ✅ Create a Task

- **Method:** `POST`
- **Path:** `/tasks`
- **Description:** Creates a new task.

#### 🔸 Request Body

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "due_date": "YYYY-MM-DDTHH:MM:SSZ",
  "status": "pending | completed | not-done"
}
```

#### 🔸 Responses

- **201 Created**

```json
{
  "id": "1",
  "title": "Sample Task",
  "description": "Test task",
  "due_date": "2025-07-18T12:00:00Z",
  "status": "pending"
}
```

- **400 Bad Request**

```json
{
  "error": "title cannot be empty"
}
```

### 2. 🔍 Get a Task by ID

- **Method:** `GET`
- **Path:** `/tasks/:id`
- **Description:** Retrieves a task by its ID.

#### 🔸 Responses

- **200 OK**

```json
{
  "id": "1",
  "title": "Sample Task",
  "description": "Test task",
  "due_date": "2025-07-18T12:00:00Z",
  "status": "pending"
}
```

- **404 Not Found**

```json
{
  "error": "task with ID 1 not found"
}
```

### 3. 📄 Get All Tasks

- **Method:** `GET`
- **Path:** `/tasks`
- **Description:** Retrieves a list of all tasks.

#### 🔸 Responses

- **200 OK**

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

- **500 Internal Server Error**

```json
{
  "error": "internal server error"
}
```

### 4. 📝 Update a Task

- **Method:** `PUT`
- **Path:** `/tasks/:id`
- **Description:** Updates an existing task.

#### 🔸 Request Body

```json
{
  "title": "string",
  "description": "string",
  "due_date": "YYYY-MM-DDTHH:MM:SSZ",
  "status": "pending | completed | not-done"
}
```

#### 🔸 Responses

- **200 OK**

```json
{
  "id": "1",
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-07-18T12:00:00Z",
  "status": "completed"
}
```

- **400 Bad Request**

```json
{
  "error": "invalid status: done"
}
```

- **404 Not Found**

```json
{
  "error": "task with ID 1 not found"
}
```

### 5. ❌ Delete a Task

- **Method:** `DELETE`
- **Path:** `/tasks/:id`
- **Description:** Deletes a task by its ID.

#### 🔸 Responses

- **204 No Content**

- **404 Not Found**

```json
{
  "error": "task with ID 1 not found"
}
```

## ⚠️ Error Handling

- All errors are returned in **JSON** format:

```json
{ "error": "description of the issue" }
```

- Common HTTP Status Codes:
  | Code | Meaning |
  |------|---------|
  | 200 | OK |
  | 201 | Created |
  | 204 | No Content (delete success) |
  | 400 | Bad Request (validation or malformed input) |
  | 404 | Not Found (resource does not exist) |
  | 500 | Internal Server Error |

## 🧪 Example Usage (cURL)

### ➕ Create a task

```bash
curl -X POST http://localhost:8080/tasks   -H "Content-Type: application/json"   -d '{"id":"1","title":"Sample Task","description":"Test task","due_date":"2025-07-18T12:00:00Z","status":"pending"}'
```

### 🔎 Get a task

```bash
curl http://localhost:8080/tasks/1
```

### 🗂️ Get all tasks

```bash
curl http://localhost:8080/tasks
```

### 🛠️ Update a task

```bash
curl -X PUT http://localhost:8080/tasks/1   -H "Content-Type: application/json"   -d '{"title":"Updated Task","description":"New info","due_date":"2025-07-20T10:00:00Z","status":"completed"}'
```

### 🗑️ Delete a task

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## ⚙️ MongoDB Configuration & Connection

This API uses MongoDB to store task data. Ensure you have MongoDB installed and running locally or use MongoDB Atlas for a cloud-hosted solution.

### 🔧 Local MongoDB Setup

1. **Install MongoDB** (if not already installed):  
   https://www.mongodb.com/try/download/community

2. **Start MongoDB Server**:

   ```bash
   mongod --dbpath <your_db_path>
   ```

3. **Default Connection String**:
   The app connects to MongoDB using the following URI by default:

   ```go
   const connectionString = "mongodb://localhost:27017"
   ```

4. **Database & Collection Used**:
   - Database: `movies`
   - Collection: `movies`

### 🛠️ Connection Code (Go)

Inside `setup.go`, the MongoDB connection is initialized as:

```go
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
client, err := mongo.Connect(context.TODO(), clientOptions)
```

And the collection is retrieved using:

```go
mongoClient.Database("movies").Collection("movies")
```

### 🌐 Using MongoDB Atlas

1. Create a free MongoDB Atlas account at https://www.mongodb.com/cloud/atlas
2. Create a new cluster and database
3. Whitelist your IP and create a user
4. Replace your connection string with the one provided by Atlas, like:

```go
const connectionString = "mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority"
```

Make sure to update your Go app accordingly.

---
