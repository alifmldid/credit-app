package server

import (
	"transaction-service/transaction"

	"github.com/gin-gonic/gin"
)

func transRoutes(r *gin.Engine, controller transaction.TransactionController){
	transaction := r.Group("/transaction")
	transaction.POST("/", controller.CreateTransaction)
}