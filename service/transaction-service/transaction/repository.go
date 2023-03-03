package transaction

import (
	"context"

	"gorm.io/gorm"
)

type TransactionRepository interface{
	Save(c context.Context, transaction Transaction) (id string, err error)
	GetTransactionById(c context.Context, id string) (transaction Transaction, err error)
}

type transactionRepository struct{
	Conn *gorm.DB
}

func NewTransactionRepository(Conn *gorm.DB) TransactionRepository{
	return &transactionRepository{Conn}
}

func (repo *transactionRepository) Save(c context.Context, transaction Transaction) (id string, err error){
	err = repo.Conn.Create(&transaction).Error
	if err != nil {
		return "", err
	}

	return transaction.TransactionID, nil	
}

func (repo *transactionRepository) GetTransactionById(c context.Context, id string) (transaction Transaction, err error){
	err = repo.Conn.Where("transaction_id = ?", id).First(&transaction).Error

	return transaction, err
}