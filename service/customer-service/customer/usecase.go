package customer

import (
	"context"
	"time"

	gonanoid "github.com/matoous/go-nanoid"
	"golang.org/x/crypto/bcrypt"
)

type CustomerUsecase interface{
	Register(c context.Context, data Customer) (custResponse CustomerResponse, err error)
	GetUser(c context.Context, id string) (custResponse CustomerResponse, err error)
	SetLimit(c context.Context, limit CustomerLimit) ([]CustomerLimit, error)
	UpdateLimit(c context.Context, custId string, tenor int, payload CreditPayload) (limit []CustomerLimit, err error)
}

type customerUsecase struct{
	customerRepository CustomerRepository	
}

func NewCustomerUsecase(customerRepository CustomerRepository) CustomerUsecase{
	return &customerUsecase{customerRepository}
}

func (uc *customerUsecase) Register(c context.Context, data Customer) (custResponse CustomerResponse, err error){
	var customer Customer

	randId, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 16)
	if err != nil {
		return CustomerResponse{}, err
	}

	customer.CustomerId = "cust-"+randId
	customer.NIK = data.NIK

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		return CustomerResponse{}, err
	}
	customer.Password = string(hashedPassword)

	customer.FullName = data.FullName
	customer.LegalName = data.LegalName
	customer.BirthPlace = data.BirthPlace
	customer.BirthDate = data.BirthDate
	customer.IDCard = data.IDCard
	customer.SelfiePhoto = data.SelfiePhoto
	customer.Status = "Not Verified"
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	id, err := uc.customerRepository.Save(c, customer)

	if err != nil {
		return CustomerResponse{}, err
	}

	customerData, err := uc.customerRepository.GetCustomerById(c, id)

	if err != nil {
		return CustomerResponse{}, err
	}
	
	custResponse.CustomerId = customer.CustomerId
	custResponse.NIK = customerData.NIK
	custResponse.FullName = customerData.FullName
	custResponse.LegalName = customerData.LegalName
	custResponse.BirthPlace = customerData.BirthPlace
	custResponse.BirthDate = customerData.BirthDate
	custResponse.IDCard = customerData.IDCard
	custResponse.SelfiePhoto = customerData.SelfiePhoto
	custResponse.Status = customerData.Status
	custResponse.CreatedAt = customerData.CreatedAt
	custResponse.UpdatedAt = customerData.UpdatedAt	

	return custResponse, err
}

func (uc *customerUsecase) GetUser(c context.Context, id string) (custResponse CustomerResponse, err error){
	customer, err := uc.customerRepository.GetCustomerById(c, id)

	custResponse.CustomerId = customer.CustomerId
	custResponse.NIK = customer.NIK
	custResponse.FullName = customer.FullName
	custResponse.LegalName = customer.LegalName
	custResponse.BirthPlace = customer.BirthPlace
	custResponse.BirthDate = customer.BirthDate
	custResponse.IDCard = customer.IDCard
	custResponse.SelfiePhoto = customer.SelfiePhoto
	custResponse.Status = customer.Status
	custResponse.CreatedAt = customer.CreatedAt
	custResponse.UpdatedAt = customer.UpdatedAt

	return custResponse, err
}

func (uc *customerUsecase) SetLimit(c context.Context, limit CustomerLimit) ([]CustomerLimit, error){
	limit.Balance = limit.Amount
	limit.CreatedAt = time.Now()
	limit.UpdatedAt = time.Now()

	err := uc.customerRepository.AddLimit(c, limit)

	if err != nil {
		return []CustomerLimit{}, err
	}

	custId := limit.CustomerID
	limitData, err := uc.customerRepository.GetLimitByCustomerId(c, custId)

	if err != nil {
		return []CustomerLimit{}, err
	}

	return limitData, err
}

func (uc *customerUsecase) UpdateLimit(c context.Context, custId string, tenor int, payload CreditPayload) (limit []CustomerLimit, err error){
	err = uc.customerRepository.UpdateLimit(c, custId, tenor, payload)

	if err != nil {
		return []CustomerLimit{}, err
	}

	limit, err = uc.customerRepository.GetLimitByCustomerId(c, custId)

	if err != nil {
		return []CustomerLimit{}, err
	}

	return limit, err
}