package config

import (
	"os"
	"strconv"
)

func GetString(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func GetInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		v2, err := strconv.Atoi(v)
		if err != nil {
			return def
		}
		return v2
	}
	return def
}

func GetInt64(key string, def int64) int64 {
	if v := os.Getenv(key); v != "" {
		v2, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return def
		}
		return v2
	}
	return def
}

func GetBool(key string, def bool) bool {
	if v := os.Getenv(key); v != "" {
		v2, err := strconv.ParseBool(v)
		if err != nil {
			return def
		}
		return v2
	}
	return def
}
