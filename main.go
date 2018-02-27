package main

import (
	"./route"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	route.Routing(router)
	router.Run(":8080")
}
