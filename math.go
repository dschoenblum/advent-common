package main

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func maxUint64(v1, v2 uint64) uint64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func divideAndRoundUp(dividend, divisor int) int {
	return (dividend + divisor - 1) / divisor
}

func divideAndRoundUpUint64(dividend, divisor uint64) uint64 {
	return (dividend + divisor - 1) / divisor
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func leastCommonMultiple(a, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}
