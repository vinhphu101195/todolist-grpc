package main

import (
	"context"
	"net"
	"todo-grpc/models"
	"todo-grpc/proto"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"google.golang.org/grpc/reflection"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type server struct {
	db *gorm.DB
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	db := models.SetupModels() // new

	srv := grpc.NewServer()

	proto.RegisterTodoModelServiceServer(srv, &server{db: db})

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	db := s.db
	//completed, userid, id, title := request.GetCompleted(), request.GetUserid(), request.GetId(), request.GetTitle()
	todo := request.GetToDo()

	result := db.NewRecord(&todo)

	return &proto.CreateResponse{KeyCreated: result}, nil
}

func (s *server) Read(ctx context.Context, request *proto.ReadRequest) (*proto.ReadResponse, error) {
	db := s.db

	var todo proto.TodoModel
	todoID := request.GetId()
	db.First(&todo, todoID)

	if todo.Id == 0 {
		return nil, status.Error(codes.Unknown, "No todo found!")

	}

	complete := false
	if todo.Completed == 0 {
		complete = false
	} else {
		complete = true
	}

	result := proto.TransformedTodo{Id: todo.Id, Title: todo.Title, Completed: complete, Userid: todo.Userid}
	return &proto.ReadResponse{ToDo: &result}, nil
}

func (s *server) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	db := s.db

	var todo proto.TodoModel
	todoID := request.GetToDo().Id
	title := request.GetToDo().Title
	completed := request.GetToDo().Completed

	db.First(&todo, todoID)

	if todo.Id == 0 {
		return nil, status.Error(codes.Unknown, "No todo found!")
	}
	db.Model(&todo).Update("title", title)
	db.Model(&todo).Update("completed", completed)

	return &proto.UpdateResponse{KeyUpdated: true}, nil

}

func (s *server) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	db := s.db

	var todo proto.TodoModel
	todoID := request.GetId()
	db.First(&todo, todoID)

	if todo.Id == 0 {
		return nil, status.Error(codes.Unknown, "No todo found!")
	}

	db.Delete(&todo)
	return &proto.DeleteResponse{KeyDeleted: true}, nil

}
