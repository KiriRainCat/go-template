package main

import (
	"github.com/gin-gonic/gin"
	"go-template/Common/DBConnection"
	"go-template/Routers"
)

func main() {
	server := gin.Default()
	DBConnection.Init()
	Routers.Init(server)
}
