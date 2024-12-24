# Go Task API

A simple task management REST API built with Go (Golang). This API allows you to create, read, update, and delete tasks. It demonstrates basic CRUD operations and can serve as a starting point for building more complex backend applications in Go.

## Features

- **Create** a task
- **Retrieve** all tasks or a specific task by ID
- **Update** a task
- **Delete** a task

## Technologies

- Go (Golang)
- Gorilla Mux Router
- In-memory data storage (for simplicity)

## Installation

### Prerequisites

Make sure you have Go installed on your machine. If you don't have Go installed, you can download it from the official Go website: [https://golang.org/dl/](https://golang.org/dl/)

### Setting Up the Project

1. Clone this repository:

   ```bash
   git clone https://github.com/Dirie/go-task-api.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-task-api
   ```

3. Install dependencies:

   This project uses the Gorilla Mux router. To install it, run:

   ```bash
   go get -u github.com/gorilla/mux
   ```

4. Run the application:

   ```bash
   go run main.go models.go
   ```

   The server will start on port `8080`. You should see the message:

   ```
   Server is running on port 8080...
   ```

## API Endpoints

### 1. **Create a Task**

- **URL**: `/tasks`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "title": "Task title",
    "description": "Task description",
    "completed": false
  }
  ```

- **Response**:
  ```json
  {
    "id": "1",
    "title": "Task title",
    "description": "Task description",
    "completed": false
  }
  ```

### 2. **Get All Tasks**

- **URL**: `/tasks`
- **Method**: `GET`

- **Response**:
  ```json
  [
    {
      "id": "1",
      "title": "Task 1",
      "description": "Description of task 1",
      "completed": false
    },
    {
      "id": "2",
      "title": "Task 2",
      "description": "Description of task 2",
      "completed": true
    }
  ]
  ```

### 3. **Get a Single Task by ID**

- **URL**: `/tasks/{id}`
- **Method**: `GET`
- **Parameters**: `id` (task ID)
  
- **Response**:
  ```json
  {
    "id": "1",
    "title": "Task title",
    "description": "Task description",
    "completed": false
  }
  ```

### 4. **Update a Task by ID**

- **URL**: `/tasks/{id}`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "title": "Updated Task title",
    "description": "Updated task description",
    "completed": true
  }
  ```

- **Response**:
  ```json
  {
    "id": "1",
    "title": "Updated Task title",
    "description": "Updated task description",
    "completed": true
  }
  ```

### 5. **Delete a Task by ID**

- **URL**: `/tasks/{id}`
- **Method**: `DELETE`
- **Parameters**: `id` (task ID)

- **Response**: `204 No Content` (No content is returned upon successful deletion)

## Example Usage

You can test the API with tools like [Postman](https://www.postman.com/) or using `curl` commands.

### Example: Create a Task using `curl`

```bash
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title": "New Task", "description": "A new task description", "completed": false}'
```

### Example: Get All Tasks using `curl`

```bash
curl http://localhost:8080/tasks
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.



