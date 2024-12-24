# Todo List API

A simple Todo List application built with Go using the Chi router, Postgres, GORM and Golang Migrate. This project provides a RESTful API for managing todos, allowing users to create, read, update, and delete todo items.

## Features

- Create a new todo
- Retrieve all todos
- Update an existing todo
- Delete a todo
- Mark todos as completed

## Technologies Used

- Go (version 1.21.3)
- Chi (v5.2.0) for routing
- JSON for data interchange
- Golang-migrate for migrations
- Postgres for database
- GORM for object relational mapping 

## Getting Started

### Prerequisites

- Go installed on your machine (version 1.21.3 or higher)
- Git for version control

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/daniyalumer/todo-list-go-chi.git
   ```

2. Navigate to the project directory:

   ```bash
   cd todo-list-go-chi
   ```

3. Install the dependencies:

   ```bash
   go mod tidy
   ```

### Running the Application

To start the server, run the following command:

```bash
go run .
```
Or
```bash
go run github.com/daniyalumer/todo-list-go-chi
```

The server will start on `http://localhost:3000`.

### API Endpoints

- **GET /todo/get**: Retrieve all todos
- **POST /todo/post**: Create a new todo
  - Body: `description=<todo_description>`
- **PUT /todo/put**: Update an existing todo
  - Body: `id=<todo_id>&description=<new_description>&completed=<true|false>`
- **DELETE /todo/delete**: Delete a todo
  - Body: `id=<todo_id>`

### Example Requests

1. **Create a Todo**

   ```bash
   curl -X POST http://localhost:3000/todo/post -d "description=Buy groceries"
   ```

2. **Get Todos**

   ```bash
   curl http://localhost:3000/todo/get
   ```

3. **Update a Todo**

   ```bash
   curl -X PUT http://localhost:3000/todo/put -d "id=1&description=Buy groceries and cook dinner&completed=true"
   ```
   
   ```bash
   http://localhost:3000/todo/put?id=2&completed=true
   ```

4. **Delete a Todo**

   ```bash
   curl -X DELETE http://localhost:3000/todo/delete -d "id=1"
   ```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.