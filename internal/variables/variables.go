package variables

import "os"

var Address string
var Port string

func init() {
	Address = getEnv("APP_ADDRESS", "0.0.0.0")
	Port = getEnv("APP_PORT", "8080")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
