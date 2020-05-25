# todolist-grpc
## to run app: run main.go
## in client folder
## in user/server folder
## in todoList/server folder 

### Listening and serving HTTP on :8080

### Register POST : http://localhost:8080/register
### login POST: http://localhost:8080/login
### logout GET: http://localhost:8080/logout

### for user profile
### GET all: http://localhost:8080/user
### GET by ID: http://localhost:8080/user/:userid
### PUT: http://localhost:8080/user/:userid
### DELETE: http://localhost:8080/user/:userid

### for todo task
### GET all: http://localhost:8080/api/v1/todos/
### GET by ID: http://localhost:8080/api/v1/todos/:id
### POST: http://localhost:8080/api/v1/todos/
### PUT: http://localhost:8080/api/v1/todos/:id
### DELETE: http://localhost:8080/api/v1/todos/:id

### --------------------------------------------------------------------------------------------------
### To create and edit the todo task, we use struct: title(string), completed (1 is true,else is false), userid(int).
### userid, and the rest will be created by gorm framework