package main

import (
	"go-template/Common/DB/MySQL"
	"go-template/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	MySQL.Init()
	Router.Init(server)
}
