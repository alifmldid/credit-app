package server

import (
	"lending/transaction"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine){
	transRepo := transaction.NewTransactionRepository()
	transUsecase := transaction.NewTransactionUsecase(transRepo)
	transController := transaction.NewTransactionController(transUsecase)

	registerTransRoute(r, transController)
}