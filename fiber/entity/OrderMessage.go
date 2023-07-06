package entity

import (
	"fiber/learning/database"
)

type OrderMessage struct {
	Id             int64               `json:"id"`
	OrderId        int64               `json:"orderId"`
	CustomerId     int64               `json:"customerId"`
	LineOfBusiness int                 `json:"lineOfBusiness"`
	Phone          string              `json:"phone"`
	OrderStatus    database.NullString `json:"orderStatus"`
	Text           database.NullString `json:"text"`
	Delivered      bool                `json:"delivered"`
	CreatedAt      database.NullTime   `json:"createdAt"`
	UpdatedAt      database.NullTime   `json:"updatedAt"`
}
