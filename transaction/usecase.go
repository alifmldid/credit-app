package transaction

import (
	"context"
	"runtime"
)

var transData = make(chan Transaction)
var custData = make(chan Customer)
var errors = make(chan error)

type TransactionUsecase interface{
	Insert(c context.Context, trans Transaction) (transResponse TransactionResponse, err error)
}

type transactionUsecase struct{
	transactionRepository TransactionRepository
}

func NewTransactionUsecase(transactionRepository TransactionRepository) TransactionUsecase{
	return &transactionUsecase{transactionRepository}
}

func (uc *transactionUsecase) Insert(c context.Context, trans Transaction) (transResponse TransactionResponse, err error){
    runtime.GOMAXPROCS(2)

	go uc.transactionRepository.Save(c, trans)

	var transData = <- transData
	var getError1 = <- errors

	if getError1 != nil {
		return TransactionResponse{}, getError1
	}

	go uc.transactionRepository.GetCustomer(c, transData.CustomerID)

	var custData = <- custData
	var getError2 = <- errors

	if getError2 != nil {
		return TransactionResponse{}, getError2
	}

	custId := trans.CustomerID
	tenor := trans.Tenor
	var payload CreditPayload
	payload.Price = trans.Price

	go uc.transactionRepository.UpdateLimit(c, custId, tenor, payload)

	var getError3 = <- errors

	if getError2 != nil {
		return TransactionResponse{}, getError3
	}

	transResponse.TransactionID = transData.TransactionID
	transResponse.Customer = custData
	transResponse.Asset = transData.Asset
	transResponse.Price = transData.Price
	transResponse.Tenor = transData.Tenor
	transResponse.Interest = transData.Interest
	transResponse.Fee = transData.Fee
	transResponse.CreatedAt = transData.CreatedAt
	transResponse.UpdatedAt = transData.UpdatedAt

	return transResponse, nil
}
