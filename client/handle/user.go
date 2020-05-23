package handle

import (
	"net/http"
	"strconv"
	"todo-grpc/user/userproto"

	"github.com/gin-gonic/gin"
)

// GetAllUser ...
func GetAllUser(c *gin.Context) {
	UserClient, connect := initUser()
	defer connect.Close()

	req := &userproto.GetAllUserRequest{}
	if response, err := UserClient.GetAllUser(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response.Users})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

//GetUserSingle ...
func GetUserSingle(c *gin.Context) {
	UserClient, connect := initUser()
	defer connect.Close()

	userID := c.Param("userid")
	req := &userproto.GetUserRequest{Userid: userID}
	if response, err := UserClient.GetUser(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response.User})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

//UpdateUser ...
func UpdateUser(c *gin.Context) {
	UserClient, connect := initUser()
	defer connect.Close()

	userID, _ := strconv.Atoi(c.Param("userid"))
	password := c.PostForm("password")
	email := c.PostForm("email")
	req := &userproto.UpdateUserRequest{User: &userproto.UserModel{Id: int32(userID), Password: password, Email: email}}
	if response, err := UserClient.UpdateUser(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": response.MessageUdated})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}

// DeleteUser ...
func DeleteUser(c *gin.Context) {
	UserClient, connect := initUser()
	defer connect.Close()

	userID, _ := strconv.Atoi(c.Param("userid"))
	req := &userproto.DeleteUserRequest{Userid: int32(userID)}
	if response, err := UserClient.DeleteUser(c, req); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": response.MessageDeleted})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
	}
}
