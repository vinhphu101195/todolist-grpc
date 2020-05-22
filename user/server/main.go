package main

import (
	"context"
	"fmt"
	"net"
	"todo-grpc/user/models"
	"todo-grpc/user/userproto"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type userServer struct {
	db *gorm.DB
}

func main() {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err)
	}
	db := models.SetupModels()
	srv := grpc.NewServer()
	userproto.RegisterUserModelServiceServer(srv, &userServer{db: db})

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *userServer) Register(ctx context.Context, request *userproto.RegisterRequest) (*userproto.RegisterResponse, error) {
	db := s.db
	user := request.GetUser()

	result := db.NewRecord(&user)
	db.Create(&user)
	return &userproto.RegisterResponse{KeyRegister: result}, nil
}

func (s *userServer) GetAllUser(ctx context.Context, request *userproto.GetAllUserRequest) (*userproto.GetAllUserResponse, error) {
	db := s.db
	_users := []*userproto.UserModel{}

	db.Find(&_users)
	if len(_users) <= 0 {
		return nil, status.Error(codes.Unknown, "No User found!")
	}

	return &userproto.GetAllUserResponse{Users: _users}, nil
}

func (s *userServer) GetUser(ctx context.Context, request *userproto.GetUserRequest) (*userproto.GetUserResponse, error) {
	db := s.db

	var user userproto.UserModel
	userID := request.GetUserid()
	db.First(&user, userID)

	if user.Id == 0 {
		return nil, status.Error(codes.Unknown, "No user found!")
	}

	return &userproto.GetUserResponse{User: &user}, nil
}

func (s *userServer) UpdateUser(ctx context.Context, request *userproto.UpdateUserRequest) (*userproto.UpdateUserResponse, error) {
	db := s.db

	var user userproto.UserModel
	userID := request.GetUser().Id
	password := request.GetUser().Password
	email := request.GetUser().Email
	db.First(&user, userID)

	if user.Id == 0 {
		return nil, status.Error(codes.Unknown, "No user found!")
	}

	fmt.Println(&user)
	fmt.Println(user.Id)
	db.Model(&user).Update("password", password)
	db.Model(&user).Update("email", email)
	fmt.Println(&user)
	return &userproto.UpdateUserResponse{MessageUdated: "User updated successfully! "}, nil

}

func (s *userServer) DeleteUser(ctx context.Context, request *userproto.DeleteUserRequest) (*userproto.DeleteUserResponse, error) {
	db := s.db

	var user userproto.UserModel
	userID := request.GetUserid()
	db.First(&user, userID)

	if user.Id == 0 {
		return nil, status.Error(codes.Unknown, "No todo found!")
	}

	db.Delete(&user)
	return &userproto.DeleteUserResponse{MessageDeleted: "Todo deleted successfully!"}, nil

}
