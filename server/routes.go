package server

import (
	"lending/transaction"

	"github.com/gin-gonic/gin"
)

func registerTransRoute(r *gin.Engine, transController transaction.TransactionController){
	user := r.Group("/transaction")
	user.POST("/", transController.TransactionInsert)
}