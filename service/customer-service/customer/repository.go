package customer

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface{
	Save(c context.Context, customer Customer) (id string, err error)
	GetCustomerByNIK(c context.Context, nik string) (customer Customer, err error)
	GetCustomerById(c context.Context, id string) (customer Customer, err error)
	AddLimit(c context.Context, limit CustomerLimit) (err error)
	GetLimitByCustomerId(c context.Context, custId string) (limit []CustomerLimit, err error)
	UpdateLimit(c context.Context, custId string, tenor int, payload CreditPayload) (err error)
}

type customerRepository struct{
	Conn *gorm.DB
}

func NewCutomerRepository(Conn *gorm.DB) CustomerRepository{
	return &customerRepository{Conn}
}

func (repo *customerRepository) Save(c context.Context, customer Customer) (id string, err error){
	err = repo.Conn.Create(&customer).Error
	if err != nil {
		return "", err
	}

	return customer.CustomerId, nil
}

func (repo *customerRepository) GetCustomerByNIK(c context.Context, nik string) (customer Customer, err error){
	err = repo.Conn.Where("nik = ?", nik).First(&customer).Error

	return customer, err
}

func (repo *customerRepository) GetCustomerById(c context.Context, id string) (customer Customer, err error){
	err = repo.Conn.Where("customer_id = ?", id).First(&customer).Error

	return customer, err
}

func (repo *customerRepository) AddLimit(c context.Context, limit CustomerLimit) (err error){
	err = repo.Conn.Create(&limit).Error
	if err != nil {
		return err
	}

	return nil	
}

func (repo *customerRepository) GetLimitByCustomerId(c context.Context, custId string) (limit []CustomerLimit, err error){
	err = repo.Conn.Where("customer_id = ?", custId).Find(&limit).Error

	return limit, err
}

func (repo *customerRepository) UpdateLimit(c context.Context, custId string, tenor int, payload CreditPayload) (err error){
	var limit CustomerLimit

	err = repo.Conn.Where("customer_id = ?", custId).Where("tenor = ?", tenor).First(&limit).Error

	if err != nil{
		return err
	}

	limit.Balance = limit.Balance-payload.Price
	limit.UpdatedAt = time.Now()

	err = repo.Conn.Where("customer_id = ?", custId).Where("tenor = ?", tenor).Save(&limit).Error

	if err != nil{
		return err
	}

	return nil
}