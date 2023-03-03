package transaction

import "time"

type Transaction struct{
	TransactionID string `json:"transaction_id"`
	CustomerID string `json:"customer_id"`
	Asset string `json:"asset"`
	Price int `json:"price"`
	Tenor int `json:"tenor"`
	Interest float32 `json:"interest"`
	Fee int `json:"fee"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}