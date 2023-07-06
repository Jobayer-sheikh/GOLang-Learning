package entity

import (
	"time"
)

type Message struct {
	Id             int64     `json:"id"`
	CustomerId     int64     `json:"customerId"`
	LineOfBusiness int       `json:"lineOfBusiness"`
	Phone          string    `json:"phone"`
	Text           string    `json:"text"`
	Delivered      bool      `json:"delivered"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
