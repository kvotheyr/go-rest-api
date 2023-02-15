package env

import (
	"os"
	"strconv"
)

func GetString(key, defaults string) string {
	envVal, found := os.LookupEnv(key)
	if !found {
		return defaults
	}

	return envVal
}

func GetInt(key string, defaults int) int {
	v, err := strconv.Atoi(os.Getenv(key))

	if err != nil {
		return defaults
	}

	return v
}
