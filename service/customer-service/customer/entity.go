package customer

import "time"

type Customer struct{
	CustomerId string
	NIK int `form:"nik"`
	Password string `form:"password"`
	FullName string `form:"full_name"`
	LegalName string `form:"legal_name"`
	BirthPlace string `form:"birth_place"`
	BirthDate string `form:"birth_date"`
	IDCard string
	SelfiePhoto string
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomerResponse struct{
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

type CustomerLimit struct{
	CustomerID string `json:"customer_id"`
	Tenor int `json:"tenor"`
	Amount int `json:"amount"`
	Balance int `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreditPayload struct{
	Price int `json:"price"`
}