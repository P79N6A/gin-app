package main

import (
	"google.golang.org/grpc"
	"log"
	"gin-app/test/grpc/hello"
	"os"
	"context"
	"time"
)

const (
	address     = "localhost:5000"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := hello.NewHelloServiceClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	userService := hello.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := client.SayHello(ctx, &hello.Request{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", result.Message)

	users, err := userService.GetUser(ctx, &hello.UserRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("user response: %v", users.Users)
}
