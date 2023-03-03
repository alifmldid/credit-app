package transaction

import (
	"context"
	"time"

	gonanoid "github.com/matoous/go-nanoid"
)

type TransactionUsecase interface{
	CreateTransaction(c context.Context, data Transaction) (Transaction, error)
}

type transactionUsecase struct{
	transactionRepository TransactionRepository
}

func NewTransactionUsecase(transactionRepository TransactionRepository) TransactionUsecase{
	return &transactionUsecase{transactionRepository}
}

func (uc *transactionUsecase) CreateTransaction(c context.Context, trans Transaction) (Transaction, error){
	randId, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 16)
	if err != nil {
		return Transaction{}, err
	}

	trans.TransactionID = "trans-"+randId
	trans.CreatedAt = time.Now()
	trans.UpdatedAt = time.Now()

	id, err := uc.transactionRepository.Save(c, trans)

	if err != nil {
		return Transaction{}, err
	}

	transData, err := uc.transactionRepository.GetTransactionById(c, id)

	if err != nil {
		return Transaction{}, err
	}

	return transData, nil
}