package customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController interface{
	CustomerRegister(c *gin.Context)
	GetUser(c *gin.Context)
	SetCustomerLimit(c *gin.Context)
	UpdateLimit(c *gin.Context)
}

type customerController struct{
	customerUsecase CustomerUsecase
}

func NewCustomerController(customerUsecase CustomerUsecase) CustomerController{
	return &customerController{customerUsecase}
}

func (controller *customerController) CustomerRegister(c *gin.Context){
	var customer Customer
	c.ShouldBind(&customer)

    idcard, err := c.FormFile("id_card")

    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "No file is received",
        })
        return
    }

    if err := c.SaveUploadedFile(idcard, "photo/idcard/" + idcard.Filename); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "message": "Unable to save the file",
        })
        return
    }

	idPath := "photo/idcard/" + idcard.Filename
	customer.IDCard = idPath

	selfie, err := c.FormFile("selfie_photo")

    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "No file is received",
        })
        return
    }

    if err := c.SaveUploadedFile(selfie, "photo/selfie/" + selfie.Filename); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "message": "Unable to save the file",
        })
        return
    }

	selfiePath := "photo/selfie/" + selfie.Filename
	customer.SelfiePhoto = selfiePath

	customerData, err := controller.customerUsecase.Register(c, customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
        return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": customerData,
	})
}

func (controller *customerController) GetUser(c *gin.Context){
	id := c.Param("id")

	customer, err := controller.customerUsecase.GetUser(c, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": customer,
	})
}

func (controller *customerController) SetCustomerLimit(c *gin.Context){
	var customerLimit CustomerLimit
	c.ShouldBind(&customerLimit)

	limit, err := controller.customerUsecase.SetLimit(c, customerLimit)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
        return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": limit,
	})
}

func (controller *customerController) UpdateLimit(c *gin.Context){
	var creditPayload CreditPayload
	c.ShouldBind(&creditPayload)

	custId := c.Param("customer_id")
	tenor, _ := strconv.Atoi(c.Param("tenor"))

	limit, err := controller.customerUsecase.UpdateLimit(c, custId, tenor, creditPayload)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
        return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": limit,
	})	
}