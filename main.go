package main

import (
	"github.com/fujiirikiya/simple-crud/route"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	route.Routing(router)
	router.Run(":8080")
}
