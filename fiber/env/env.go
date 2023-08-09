package env

import (
	"os"
)

const (
	Local               = "Local"
	Development         = "Development"
	Production          = "Production"
	EnvironmentVariable = "GO_ENVIRONMENT"
)

func GetEnvironment(src string) string {
	return os.Getenv(src)
}

func IsDevelopment() bool {
	return os.Getenv(EnvironmentVariable) == Development
}

func IsProduction() bool {
	return os.Getenv(EnvironmentVariable) == Production
}

func IsLocalDevelopment() bool {
	return os.Getenv(EnvironmentVariable) == Local
}

func GetApplicationPort() string {
	return os.Getenv("PORT")
}
