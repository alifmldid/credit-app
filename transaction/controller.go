package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController interface{
	TransactionInsert(c *gin.Context)
}

type transactionController struct{
	transactionUsecase TransactionUsecase
}

func NewTransactionController(transactionUsecase TransactionUsecase)TransactionController{
	return &transactionController{transactionUsecase}
}

func (controller *transactionController) TransactionInsert(c *gin.Context){
	var payload Transaction
	c.ShouldBindJSON(&payload)

	trans, err := controller.transactionUsecase.Insert(c, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": trans,
	})	
}