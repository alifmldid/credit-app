package transaction

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

var customerURL = "http://localhost:8020/customer"
var transactionURL = "http://localhost:8030/transaction"

type TransactionRepository interface{
	Save(c context.Context, trans Transaction)
	GetCustomer(c context.Context, id string)
	UpdateLimit(c context.Context, custId string, tenor int, payload CreditPayload)
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository{
	return &transactionRepository{}
}

func (repo *transactionRepository) Save(c context.Context, trans Transaction){
	var responseData NewTransactionResponse
	client := &http.Client{}
	userByte, err := json.Marshal(trans)
	if err != nil {
		transData <- Transaction{}
		errors <- err
		return
	}

	request, err := http.NewRequest("POST", transactionURL, bytes.NewBuffer(userByte))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		transData <- Transaction{}
		errors <- err
		return
	}

	response, err := client.Do(request)
	if err != nil {
		transData <- Transaction{}
		errors <- err
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		transData <- Transaction{}
		errors <- err
		return
	}

	transData <- responseData.Data
	errors <- nil
}

func (repo *transactionRepository) GetCustomer(c context.Context, id string) {	
	var responseData CustomerResponse
	client := &http.Client{}

	request, err := http.NewRequest("GET", customerURL+"/"+id, nil)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		custData <- Customer{}
		errors <- err
		return
	}

	response, err := client.Do(request)
	if err != nil {
		custData <- Customer{}
		errors <- err
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		custData <- Customer{}
		errors <- err
		return
	}

	custData <- responseData.Data
	errors <- nil
}

func (repo *transactionRepository) UpdateLimit(c context.Context, custId string, tenor int, payload CreditPayload) {
	var responseData NewTransactionResponse
	client := &http.Client{}
	userByte, err := json.Marshal(payload)
	if err != nil {
		errors <- err
	}

	tenorString := strconv.Itoa(tenor)

	request, err := http.NewRequest("PATCH", customerURL+"/limit/"+custId+"/"+tenorString, bytes.NewBuffer(userByte))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		errors <- err
	}

	response, err := client.Do(request)
	if err != nil {
		errors <- err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		errors <- err
	}

	errors <- nil
}