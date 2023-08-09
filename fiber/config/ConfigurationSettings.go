package config

type ShortMessageConfig struct {
	ApiUrl   string
	ApiKey   string
	SenderId string
}

type NotificationConfig struct {
	ApiUrl    string
	Serverkey string
	SenderId  string
}

type ConfigurationSettings struct {
	ShortMessageConfig   ShortMessageConfig
	NotificationConfig   NotificationConfig
	IsCacheEnable        bool
	IsUniversalOtpEnable bool
}
