package ordermessage

type OrderMessageQuery struct {
	Page           int   `json:"page"`
	Limit          int   `json:"limit"`
	OrderId        int64 `json:"orderId"`
	CustomerId     int64 `json:"customerId"`
	LineOfBusiness int   `json:"lineOfBusiness"`
}
