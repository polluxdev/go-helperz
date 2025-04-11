package helperz

import (
	"os"
	"strconv"
	"strings"
)

type EnvHelper interface {
	// Single
	GetEnvString(key, fallback string) string
	GetEnvInt(key string, fallback int) int
	GetEnvBool(key string, fallback bool) bool

	// Slice
	GetEnvSliceString(key string, fallback []string) []string
	GetEnvSliceInt(key string, fallback []int) []int
	GetEnvSliceBool(key string, fallback []bool) []bool
}

func GetEnvString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return fallback
}

func GetEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return fallback
}

func GetEnvSliceString(key string, fallback []string) []string {
	if value, ok := os.LookupEnv(key); ok {
		return strings.Split(value, ",")
	}
	return fallback
}

func GetEnvSliceInt(key string, fallback []int) []int {
	if value, ok := os.LookupEnv(key); ok {
		slice := strings.Split(value, ",")

		result := make([]int, 0)
		for _, item := range slice {
			if parsed, err := strconv.Atoi(item); err != nil {
				return fallback
			} else {
				result = append(result, parsed)
			}
		}

		return result
	}
	return fallback
}

func GetEnvSliceBool(key string, fallback []bool) []bool {
	if value, ok := os.LookupEnv(key); ok {
		slice := strings.Split(value, ",")

		result := make([]bool, 0)
		for _, item := range slice {
			if parsed, err := strconv.ParseBool(item); err != nil {
				return fallback
			} else {
				result = append(result, parsed)
			}
		}

		return result
	}
	return fallback
}
