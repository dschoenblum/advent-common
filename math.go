package common

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func Max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func MaxUint64(v1, v2 uint64) uint64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func DivideAndRoundUp(dividend, divisor int) int {
	return (dividend + divisor - 1) / divisor
}

func DivideAndRoundUpUint64(dividend, divisor uint64) uint64 {
	return (dividend + divisor - 1) / divisor
}

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LeastCommonMultiple(a, b int) int {
	return a * b / GreatestCommonDivisor(a, b)
}
