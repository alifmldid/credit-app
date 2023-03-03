package main

import (
	"customer-service/server"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	server.RegisterAPIService(r)
	
	r.Run(":8020")
}