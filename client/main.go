package main

import (
	"log"
	"net/http"
	"strconv"
	"todo-grpc/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewTodoModelServiceClient(conn)
	r := gin.Default()

	v1 := r.Group("/api/v1/todos")
	{
		v1.POST("/", func(c *gin.Context) {
			completed, _ := strconv.Atoi(c.PostForm("completed"))
			userid, _ := strconv.Atoi(c.PostForm("userid"))
			req := &proto.CreateRequest{ToDo: &proto.TodoModel{Title: c.PostForm("title"), Completed: int32(completed), Userid: int32(userid)}}
			if response, err := client.Create(c, req); err == nil {
				c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "response": response.KeyCreated})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})

			}
		})
		//v1.GET("/", controller.FetchAllTodo)
		v1.GET("/:id", func(c *gin.Context) {
			todoID := c.Param("id")
			req := &proto.ReadRequest{Id: todoID}
			if response, err := client.Read(c, req); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response.ToDo})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
		v1.PUT("/:id", func(c *gin.Context) {
			todoID, _ := strconv.Atoi(c.Param("id"))
			completed, _ := strconv.Atoi(c.PostForm("completed"))
			req := &proto.UpdateRequest{ToDo: &proto.TodoModel{Id: int32(todoID), Title: c.PostForm("title"), Completed: int32(completed)}}
			if _, err := client.Update(c, req); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
			} else {
				c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
			}
		})
		//v1.DELETE("/:id", controller.DeleteTodo)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
