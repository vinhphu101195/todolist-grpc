# todolist-grpc
## to run app: run main.go in client folder, and run main.go in server folder by go run main.go

### Listening and serving HTTP on :8080
### GET all: http://localhost:8080/api/v1/todos/
### GET by ID: http://localhost:8080/api/v1/todos/:id
### POST: http://localhost:8080/api/v1/todos/
### PUT: http://localhost:8080/api/v1/todos/:id
### DELETE: http://localhost:8080/api/v1/todos/:id

### --------------------------------------------------------------------------------------------------
### To create and edit the todo task, we use struct: title(string), completed (1 is true,else is false), userid(int).
### userid, and the rest will be created by gorm framework