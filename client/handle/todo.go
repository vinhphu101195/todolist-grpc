package handle

import (
	"net/http"
	"strconv"
	"todo-grpc/todoList/proto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func initTodo() (proto.TodoModelServiceClient, *grpc.ClientConn) {
	// connect todoList
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	todoClient := proto.NewTodoModelServiceClient(conn)

	return todoClient, conn
}

//CreteTodo ...
func CreteTodo(c *gin.Context) {
	todoClient, connect := initTodo()
	defer connect.Close()
	session := sessions.Default(c)

	completed, _ := strconv.Atoi(c.PostForm("completed"))
	userid := session.Get("AuthUserId")
	if userid == nil {
		c.JSON(http.StatusCreated, gin.H{"error": "unknow user"})
		return
	}

	req := &proto.CreateRequest{ToDo: &proto.TodoModel{Title: c.PostForm("title"), Completed: int32(completed), Userid: userid.(int32)}}
	if response, err := todoClient.Create(c, req); err == nil {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "response": response.KeyCreated})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

// GetAllTodo ...
func GetAllTodo(c *gin.Context) {
	todoClient, connect := initTodo()
	defer connect.Close()

	req := &proto.ReadAllRequest{}
	if response, err := todoClient.ReadAll(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response.ToDos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

//GetSingleTodo ...
func GetSingleTodo(c *gin.Context) {
	todoClient, connect := initTodo()
	defer connect.Close()

	todoID := c.Param("id")
	req := &proto.ReadRequest{Id: todoID}
	if response, err := todoClient.Read(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response.ToDo})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

//UpdateTodo ...
func UpdateTodo(c *gin.Context) {
	todoClient, connect := initTodo()
	defer connect.Close()

	todoID, _ := strconv.Atoi(c.Param("id"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	req := &proto.UpdateRequest{ToDo: &proto.TodoModel{Id: int32(todoID), Title: c.PostForm("title"), Completed: int32(completed)}}
	if response, err := todoClient.Update(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": response.MessageUdated})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

//DeleteTodo ...
func DeleteTodo(c *gin.Context) {
	todoClient, connect := initTodo()
	defer connect.Close()

	todoID, _ := strconv.Atoi(c.Param("id"))
	req := &proto.DeleteRequest{Id: int32(todoID)}
	if response, err := todoClient.Delete(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": response.MessageDeleted})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}
