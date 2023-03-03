package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController interface{
	CreateTransaction(c *gin.Context)
}

type transactionController struct{
	transactionUsecase TransactionUsecase
}

func NewTransactionController(transactionUsecase TransactionUsecase) TransactionController{
	return &transactionController{transactionUsecase}
}

func (controller *transactionController)  CreateTransaction(c *gin.Context){
	var transaction Transaction
	c.ShouldBindJSON(&transaction)

	trans, err := controller.transactionUsecase.CreateTransaction(c, transaction)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
        return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": trans,
	})
}