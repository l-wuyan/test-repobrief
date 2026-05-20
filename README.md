# Task API

A simple in-memory task management API written in Go. This service provides RESTful endpoints to create, list, and toggle the completion status of tasks. It's designed as a lightweight demonstration of a concurrent-safe HTTP server.

## Architecture

The application follows a simple client-server model with an in-memory data store. The diagram below illustrates the core components and their interactions.

```mermaid
graph TD
    Client[HTTP Client] --> Router[HTTP Router]
    Router -->|GET /tasks| HandlerGet[List Handler]
    Router -->|POST /tasks| HandlerPost[Create Handler]
    Router -->|POST /tasks/toggle?id=...| HandlerToggle[Toggle Handler]
    
    HandlerGet --> Store[TaskStore]
    HandlerPost --> Store
    HandlerToggle --> Store
    
    Store -->|sync.RWMutex| Mutex[Concurrency Control]
    Store -->|map[string]*Task| Memory[In-Memory Storage]
```

## Key Features

*   **RESTful API**: Provides standard HTTP endpoints for task management.
*   **In-Memory Storage**: Uses a Go map for fast data access (data is lost on server restart).
*   **Concurrency Safe**: Implements `sync.RWMutex` to handle multiple simultaneous requests safely.
*   **JSON Communication**: All request and response bodies use JSON format.
*   **Auto-generated IDs**: Tasks are assigned unique sequential IDs (e.g., `task-1`, `task-2`).

## Quick Start

### Prerequisites
*   Go 1.22 or later installed on your system.

### Installation & Running

1.  Clone the repository and navigate into the project directory.
    ```bash
    git clone <repository-url>
    cd test-repobrief
    ```
2.  Run the application.
    ```bash
    go run main.go
    ```
3.  The server will start and listen on port `8090`. You should see the log:
    ```
    Task API running on :8090
    ```

## Usage Examples

You can interact with the API using any HTTP client like `curl`, `wget`, or tools like Postman.

### 1. Create a New Task
**Request:**
```bash
curl -X POST http://localhost:8090/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy groceries"}'
```
**Response (HTTP 201 Created):**
```json
{
  "id": "task-1",
  "title": "Buy groceries",
  "done": false,
  "created_at": "2026-04-18T10:30:00Z"
}
```

### 2. List All Tasks
**Request:**
```bash
curl -X GET http://localhost:8090/tasks
```
**Response:**
```json
[
  {
    "id": "task-1",
    "title": "Buy groceries",
    "done": false,
    "created_at": "2026-04-18T10:30:00Z"
  }
]
```

### 3. Toggle a Task's Completion Status
**Request:**
```bash
curl -X POST "http://localhost:8090/tasks/toggle?id=task-1"
```
**Response:**
```json
{
  "id": "task-1",
  "title": "Buy groceries",
  "done": true,
  "created_at": "2026-04-18T10:30:00Z"
}
```

## Configuration

The application currently has minimal configuration:
*   **Port**: The HTTP server is hardcoded to run on port `8090`. To change this, modify the `http.ListenAndServe(":8090", nil)` line in `main.go`.
*   **Storage**: The `TaskStore` is in-memory only. There is no persistence layer or external database configuration.

## Project Structure

```
.
├── go.mod          # Go module definition
├── main.go         # Application entry point, HTTP server, and core logic
└── README.md       # This documentation file
```

**File Breakdown:**
*   `main.go`: Contains all application code:
    *   `Task` struct: The data model.
    *   `TaskStore` struct: The thread-safe in-memory store with methods for `Add`, `List`, and `Toggle`.
    *   `main()`: Sets up HTTP route handlers (`/tasks` and `/tasks/toggle`) and starts the server.

## Contributing

This is a test/demonstration project. Contributions are not expected. If you wish to fork or modify the code for your own purposes, you are free to do so.

## License

No license file was detected in the provided codebase. The code is shared without an explicit license, which typically means all rights are reserved by the author. Please contact the repository owner for licensing information.