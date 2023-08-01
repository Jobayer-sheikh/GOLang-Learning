package smsservice

var (
	SenderId string = "8809677399956"
	Type     string = "text/unicode"
	ApiKey   string = "C3000010644cf4bc808c37.91444038"
	Url      string = "https://sms.bracnet.net/smsapi"
)

type SmsCommand struct {
	ApiKey   string `json:"api_key"`
	Type     string `json:"type"`
	Contacts string `json:"contacts"`
	Msg      string `json:"msg"`
	Label    string `json:"label"`
	SenderId string `json:"senderid"`
}
