package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"hey": "heye",
		},
	)
}
