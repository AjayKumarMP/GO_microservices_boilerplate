package authservice

import (
	"dateme_server/services/authService/proto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler
func Handler(router *gin.Engine, client proto.AuthServiceClient) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("register", func(ctx *gin.Context) {
			request := &proto.RegisterRequest{
				Name:        "Ajay",
				Email:       "ajay@gmail.com",
				Password:    "test",
				PhoneNumber: int64(23232333),
			}
			loginReq := &proto.LoginRequest{
				Email:    "ajay@gmail.com",
				Password: "test",
			}
			go func() {
				if response, err := client.Login(ctx, loginReq); err != nil {
					fmt.Println("error occured", err)
				} else {
					fmt.Println(response, "success called login")
				}
			}()
			if response, err := client.Register(ctx, request); err == nil {
				ctx.JSON(
					http.StatusOK,
					gin.H{
						"hey": response,
					},
				)
			} else {
				ctx.JSON(
					http.StatusInternalServerError,
					gin.H{
						"Message": "Server error",
					},
				)
			}
		})
	}
}
