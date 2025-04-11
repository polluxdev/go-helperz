package helperz

type MathHelper interface {
	Percent(a, b float64) float64
}

func Percent(a, b float64) float64 {
	if b == 0 {
		return a * 100
	}

	result := (a - b) / b * 100
	if result < 0 {
		result = result * -1
	}

	return result
}
