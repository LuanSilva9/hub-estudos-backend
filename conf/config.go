package conf

import (
	"os"

	"github.com/joho/godotenv"
)

type GlobalConfig struct {
	Env         string
	Port        string
	DatabaseDsn string
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func LoadGlobalConfig() (*GlobalConfig, error) {
	if err := godotenv.Overload(); err != nil {
		return nil, err
	}

	return &GlobalConfig{
		Env:         getEnv("ENV"),
		Port:        getEnv("PORT"),
		DatabaseDsn: getEnv("DATABASE_DSN"),
	}, nil
}
