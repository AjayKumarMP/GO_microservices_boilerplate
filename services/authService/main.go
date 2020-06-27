package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"dateme_server/services/authService/proto"
)

type authServer struct{}

type user struct {
	name        string
	email       string
	phoneNumber int64
	password    string
}

var users []user

func main() {
	fmt.Println("starting Auth Service")
	listner, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, &authServer{})
	reflection.Register(server)

	if e := server.Serve(listner); e != nil {
		fmt.Println("Auth server is running @ 4040")
		panic(e)
	}

}

func (s *authServer) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	name, email, pswd := request.GetName(), request.GetEmail(), request.GetPassword()
	pno := request.GetPhoneNumber()
	users = append(users, user{name: name, email: email, phoneNumber: pno, password: pswd})
	result := "registered suucessfully"

	return &proto.RegisterResponse{Success: true, Message: result}, nil
}

func (s *authServer) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	email, password := request.GetEmail(), request.GetPassword()
	for _, v := range users {
		if v.email == email && v.password == password {
			return &proto.LoginResponse{Success: true, Message: "token"}, nil
		}
	}
	return &proto.LoginResponse{Success: false, Message: "Please check ur credentials"}, nil
}
