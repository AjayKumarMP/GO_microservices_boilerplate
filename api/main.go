package main

import (
	"dateme_server/api/authservice"
	"dateme_server/services/authService/proto"

	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAuthServiceClient(clientConn)

	router := gin.Default()

	authservice.Handler(router, client)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("error while ruunig server", err)
	} else {
		fmt.Println("Server up and Running")
	}

}
