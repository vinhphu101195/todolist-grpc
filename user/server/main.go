package main

import (
	"context"
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
