package helperz

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

type DataHelper interface {
	// Set default value
	SetDefaultString(value, defaultValue string) string
	SetDefaultInt(value, defaultValue int) int

	// Check if slice contains value
	ContainString(slice []string, item string) bool
	ContainStringWithEqualFold(slice []string, item string) bool
	ContainInt(slice []int, item int) bool

	// Generate data
	GenerateTimeDuration(n int, d time.Duration) time.Duration
	GenerateRandomString(length int, charset string) (string, error)
	GenerateRefID(length int, prefix string) (string, error)

	// Remove duplicate data
	RemoveDuplicateString(input []string) []string
	RemoveDuplicateInt(input []int) []int
}

func SetDefaultString(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

func SetDefaultInt(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}

	return value
}

func ContainString(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func ContainStringWithEqualFold(slice []string, item string) bool {
	for _, v := range slice {
		if strings.EqualFold(v, item) {
			return true
		}
	}
	return false
}

func ContainInt(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func GenerateTimeDuration(n int, d time.Duration) time.Duration {
	return time.Duration(n) * d
}

func GenerateRandomString(length int, charset string) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	if charset == "" {
		charset = ALPHANUMERIC_CHARSET
	}

	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}

func GenerateRefID(length int, prefix string) (string, error) {
	random, err := GenerateRandomString(length, UPPER_ALPHANUMERIC_CHARSET)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", prefix, random), nil
}

func RemoveDuplicateString(input []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

func RemoveDuplicateInt(input []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}
