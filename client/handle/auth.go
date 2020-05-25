package handle

import (
	"net/http"
	"todo-grpc/user/userproto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

//initUser ...
func initUser() (userproto.UserModelServiceClient, *grpc.ClientConn) {
	// connect User
	connUser, errUser := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if errUser != nil {
		panic(errUser)
	}
	userClient := userproto.NewUserModelServiceClient(connUser)

	return userClient, connUser
}

// Login ...
func Login(c *gin.Context) {
	UserClient, connect := initUser()
	defer connect.Close()

	session := sessions.Default(c)

	if session.Get("Auth") == true {
		c.JSON(http.StatusConflict, gin.H{"error": "Already login!!!"})
		return
	}

	userName := c.PostForm("username")
	userPassword := c.PostForm("password")

	req := &userproto.LoginRequest{Username: userName, Password: userPassword}
	if response, err := UserClient.Login(c, req); err == nil {
		session.Set("Auth", true)
		session.Set("AuthUserId", response.UserId)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"status": "login successed", "message": response.Success, "UserId": response.UserId})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

// Register ...
func Register(c *gin.Context) {
	UserClient, connect := initUser()
	defer connect.Close()

	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	req := &userproto.RegisterRequest{User: &userproto.UserModel{Username: username, Password: password, Email: email}}
	if response, err := UserClient.Register(c, req); err == nil {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User Register successfully!", "response": response.KeyRegister})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

// Logout ...
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("Auth", false)
	session.Set("AuthUserName", nil)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "logout successed"})
}
