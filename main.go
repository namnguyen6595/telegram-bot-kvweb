package main

import (
	"github.com/gin-gonic/gin"
	"kvweb-bot/handlers"
)

func main() {

	getListTransaction := &handlers.GetListTransaction{}

	getMessagehandler := &handlers.GetMessageHandler{}
	r := gin.Default()

	router := r.Group("/api")

	router.POST("/telegram", getMessagehandler.NewServe)
	router.GET("/transactions", getListTransaction.NewServe)
	r.Run(":8080")
}
