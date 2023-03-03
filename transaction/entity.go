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

type NewTransactionResponse struct{
	Message string `json:"message"`
	Data Transaction `json:"data"`
}

type CustomerResponse struct{
	Message string `json:"message"`
	Data Customer `json:"data"`
}

type TransactionResponse struct{
	TransactionID string `json:"transaction_id"`
	Customer Customer `json:"customer"`
	Asset string `json:"asset"`
	Price int `json:"price"`
	Tenor int `json:"tenor"`
	Interest float32 `json:"interest"`
	Fee int `json:"fee"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Customer struct{
	CustomerId string `json:"customer_id"`
	NIK int `json:"nik"`
	FullName string `json:"full_name"`
	LegalName string `json:"legal_name"`
	BirthPlace string `json:"birth_place"`
	BirthDate string `json:"birth_date"`
	IDCard string `json:"id_card"`
	SelfiePhoto string `json:"selfie_photo"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreditPayload struct{
	Price int `json:"price"`
}