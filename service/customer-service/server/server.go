package server

import (
	"customer-service/config"
	"customer-service/customer"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine){
	db := config.GetDBConnection()
	
	customerRepo := customer.NewCutomerRepository(db)
	customerUsecase := customer.NewCustomerUsecase(customerRepo)
	customerController := customer.NewCustomerController(customerUsecase)

	customerRoutes(r, customerController)
}