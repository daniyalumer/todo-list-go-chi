# Todo List API

A simple Todo List application built with Go using the Chi router, PostgreSQL, GORM and Golang Migrate. This project provides a RESTful API for managing todos, allowing users to create, read, update, and delete todo items.

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
- PostgreSQL for database
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
or
```bash
go run github.com/daniyalumer/todo-list-go-chi
```

The server will start on `http://localhost:3000`.

### API Endpoints

- **GET api/home/**: Print Welcome Message     
&nbsp;
- **GET api/todo/**: Retrieve all todos
- **POST api/todo/{user_id}**: Create a new todo
  - Body: `description=<todo_description>`
- **PUT api/todo/{todo_id}**: Update an existing todo
  - Body: `description=<new_description>&completed=<true|false>`
- **DELETE api/todo/{user_id}**: Delete a todo
&nbsp;
- **GET api/user/**: Retrieve all users
- **POST api/user/**: Create a new users
  - Body: `username=<username>`
- **DELETE api/user/{user_id}**: Delete a user and his todos   

### Example Requests

1. **Create a User**

   ```bash
   curl -X POST http://localhost:3000/api/user/ -d "username=xyz"
   ```

2. **Get User**

   ```bash
   curl http://localhost:3000/api/user/
   ```

3. **Delete a User**

   ```bash
   curl -X DELETE http://localhost:3000/api/user/{user}  
   ```

4. **Create a Todo**

   ```bash
   curl -X POST http://localhost:3000/api/todo/{user_id} -d "description=Buy groceries"
   ```

5. **Get Todos**

   ```bash
   curl http://localhost:3000/api/todo/
   ```

6. **Update a Todo**

   ```bash
   curl -X PUT http://localhost:3000/api/todo/{todo_id} -d "description=Buy groceries and cook dinner&completed=true"
   ```

7. **Delete a Todo**

   ```bash
   curl -X DELETE http://localhost:3000/api/todo/delete/{todo_id}  
   ```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License

This project is Open Source