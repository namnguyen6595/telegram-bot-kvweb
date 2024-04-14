package main

import (
	"github.com/gin-gonic/gin"
	"kvweb-bot/database"
	"kvweb-bot/handlers"
	"kvweb-bot/models"
)

func main() {
	db := database.Initial()
	balanceDmsClient := &models.BalanceDmsClient{
		Db: db,
	}
	topupBalanceHandler := &handlers.TopupBalanceHandler{
		balanceDmsClient,
	}
	getListTransaction := handlers.GetListTransaction{
		BalanceDmsClient: balanceDmsClient,
	}
	r := gin.Default()

	router := r.Group("/api")

	router.POST("/topup", topupBalanceHandler.NewServe)
	router.GET("/transaction", getListTransaction.NewServe)

	r.Run(":8080")
}
