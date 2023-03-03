package server

import (
	"transaction-service/config"
	"transaction-service/transaction"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine){
	db := config.GetDBConnection()
	
	transRepo := transaction.NewTransactionRepository(db)
	transUsecase := transaction.NewTransactionUsecase(transRepo)
	transController := transaction.NewTransactionController(transUsecase)

	transRoutes(r, transController)
}