package main

import (
	"log"
	"net/http"
	"strconv"
	"todo-grpc/todoList/proto"
	"todo-grpc/user/userproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// connect todoList
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	todoClient := proto.NewTodoModelServiceClient(conn)

	// connect User
	connUser, errUser := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if errUser != nil {
		panic(err)
	}
	UserClient := userproto.NewUserModelServiceClient(connUser)

	r := gin.Default()

	v1 := r.Group("/api/v1/todos")
	{
		v1.POST("/", func(c *gin.Context) {
			completed, _ := strconv.Atoi(c.PostForm("completed"))
			userid, _ := strconv.Atoi(c.PostForm("userid"))
			req := &proto.CreateRequest{ToDo: &proto.TodoModel{Title: c.PostForm("title"), Completed: int32(completed), Userid: int32(userid)}}
			if response, err := todoClient.Create(c, req); err == nil {
				c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "response": response.KeyCreated})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
		v1.GET("/", func(c *gin.Context) {
			req := &proto.ReadAllRequest{}
			if response, err := todoClient.ReadAll(c, req); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response.ToDos})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
		v1.GET("/:id", func(c *gin.Context) {
			todoID := c.Param("id")
			req := &proto.ReadRequest{Id: todoID}
			if response, err := todoClient.Read(c, req); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response.ToDo})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
		v1.PUT("/:id", func(c *gin.Context) {
			todoID, _ := strconv.Atoi(c.Param("id"))
			completed, _ := strconv.Atoi(c.PostForm("completed"))
			req := &proto.UpdateRequest{ToDo: &proto.TodoModel{Id: int32(todoID), Title: c.PostForm("title"), Completed: int32(completed)}}
			if response, err := todoClient.Update(c, req); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": response.MessageUdated})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
		v1.DELETE("/:id", func(c *gin.Context) {
			todoID, _ := strconv.Atoi(c.Param("id"))
			req := &proto.DeleteRequest{Id: int32(todoID)}
			if response, err := todoClient.Delete(c, req); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": response.MessageDeleted})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
	}

	user := r.Group("/user")
	{
		user.POST("/", func(c *gin.Context) {
			username := c.PostForm("username")
			password := c.PostForm("password")
			email := c.PostForm("email")
			req := &userproto.RegisterRequest{User: &userproto.UserModel{Username: username, Password: password, Email: email}}
			if response, err := UserClient.Register(c, req); err == nil {
				c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User Register successfully!", "response": response.KeyRegister})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
		user.GET("/", func(c *gin.Context) {
			req := &userproto.GetAllUserRequest{}
			if response, err := UserClient.GetAllUser(c, req); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response.Users})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
