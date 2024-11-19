package env

import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	ival, err := strconv.Atoi(val)
	if err != nil {
		return ival
	}
	return fallback
}
