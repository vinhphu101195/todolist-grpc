package main

import (
	"log"
	"net/http"
	"todo-grpc/client/handle"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Authorize ...
func Authorize(ctx *gin.Context) {
	session := sessions.Default(ctx)

	if session.Get("Auth") != true {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauth"})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", handle.Login)
	r.POST("/register", handle.Register)
	r.GET("/logout", Authorize, handle.Logout)

	v1 := r.Group("/api/v1/todos")
	v1.Use(Authorize)
	{
		v1.POST("/", handle.CreteTodo)
		v1.GET("/", handle.GetAllTodo)
		v1.GET("/:id", handle.GetSingleTodo)
		v1.PUT("/:id", handle.UpdateTodo)
		v1.DELETE("/:id", handle.DeleteTodo)
	}

	user := r.Group("/user")
	user.Use(Authorize)
	{
		user.GET("/", handle.GetAllUser)
		user.GET("/:userid", handle.GetUserSingle)
		user.PUT("/:userid", handle.UpdateUser)
		user.DELETE("/:userid", handle.DeleteUser)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
