package entity

import (
	"time"
)

type OrderNotification struct {
	Id             int64     `json:"id"`
	OrderId        int64     `json:"orderId"`
	CustomerId     int64     `json:"customerId"`
	LineOfBusiness int       `json:"lineOfBusiness"`
	Phone          string    `json:"phone"`
	AppToken       string    `json:"appToken"`
	Title          string    `json:"title"`
	Message        string    `json:"message"`
	Delivered      bool      `json:"delivered"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
