package main

import (
	"context"
	"gin-app/test/grpc/hello"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

const PORT string = ":5000"

type Server struct {
}

func (server *Server) SayHello(ctx context.Context, in *hello.Request) (*hello.Response, error) {
	log.Printf("receive name is: %s", in.Name)
	return &hello.Response{Message: "hello " + in.Name}, nil
}

//host和端口
const (
	host     string = "127.0.0.1"
	port            = 3306
	database        = "test"
	userName string = "root"
	password string = "123456"
)

//数据库连接串
type UserService struct {
}


func getUser(id int) hello.User {
	//打开连接
	var jdbcUrl string = userName + ":" + password + "@(" + host + ":" + strconv.Itoa(port) + ")/" + database + "?charset=utf8"

	db, err:= sql.Open("mysql", jdbcUrl)
	if err != nil {
		fmt.Println(err)
	}
	//最后关系连接
	defer db.Close()

	//查询数据
	stmt, err := db.Prepare("select * from user where id=?")
	if err != nil {
		fmt.Println(err)
	}
	var user hello.User
	fmt.Println("id = ", id)
	stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	return user
}

func (this *UserService) GetUser(ctx context.Context, request *hello.UserRequest) (*hello.UserResponse, error) {
	log.Printf("receive user request: %v", request)
	user := getUser(int(request.Id))
	fmt.Println(user)
	return &hello.UserResponse{Users: []*hello.User{&user}}, nil
}

func (this *UserService) GetUsers(ctx context.Context, request *hello.UserRequest) (*hello.UserResponse, error) {
	log.Printf("receive user request: %v", request)
	return &hello.UserResponse{Users: nil}, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, &Server{})
	hello.RegisterUserServiceServer(server, &UserService{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
