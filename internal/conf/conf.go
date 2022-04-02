package conf

import "os"

const (
	addrEnvKey  = "ADDR"
	defaultAddr = ":8080"
)

type ApplicationConfig struct {
	Addr string
}

func ParseConfig() ApplicationConfig {
	addr := getEnv(addrEnvKey, defaultAddr)

	return ApplicationConfig{
		Addr: addr,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
