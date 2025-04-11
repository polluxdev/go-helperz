package helperz

import "time"

type ConverterHelper interface {
	// string and pointer
	StringToPointer(s string) *string
	PointerToString(s *string) string

	// int and pointer
	IntToPointer(i int) *int
	PointerToInt(i *int) int

	// time and pointer
	TimeToPointer(t time.Time) *time.Time
	PointerToTime(t *time.Time) time.Time

	// bool and pointer
	BoolToPointer(b bool) *bool
	PointerToBool(b *bool) bool
}

func StringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func PointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IntToPointer(i int) *int {
	return &i
}

func PointerToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func TimeToPointer(t time.Time) *time.Time {
	return &t
}

func PointerToTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func BoolToPointer(b bool) *bool {
	return &b
}

func PointerToBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
