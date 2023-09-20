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

// GcdExtended is based on https://www.geeksforgeeks.org/multiplicative-inverse-under-modulo-m/
func GcdExtended(a, b int) (x, y, gcd int) {
	// base case
	if a == 0 {
		return 0, 1, b
	}

	// call recursively
	var x1, y1 int
	x1, y1, gcd = GcdExtended(b%a, a)

	// set x and y
	x = y1 - (b/a)*x1
	y = x1
	return
}

// ModInverse is based on https://www.geeksforgeeks.org/multiplicative-inverse-under-modulo-m/
func ModInverse(a, m int) (int, bool) {
	x, _, gcd := GcdExtended(a, m)
	if gcd != 1 {
		return 0, false
	}
	return (x%m + m) % m, true
}

// ChineseRemainder is based on https://www.geeksforgeeks.org/implementation-of-chinese-remainder-theorem-inverse-modulo-based-implementation/#
func ChineseRemainder(num, rem []int) int {
	if len(num) != len(rem) {
		panic("array lengths do not match")
	}

	product := 1
	for _, n := range num {
		product *= n
	}

	x := 0
	for i := 0; i < len(num); i++ {
		pp := product / num[i]
		inv, _ := ModInverse(pp, num[i])
		x += rem[i] * pp * inv
	}

	return x % product
}
