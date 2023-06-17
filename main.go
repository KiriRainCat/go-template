package main

import (
	"bingchat-gpt4-backend/Common/DBConnection"
	"bingchat-gpt4-backend/Routers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	DBConnection.Init()
	Routers.Init(server)
}
