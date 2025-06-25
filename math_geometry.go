package main

// 202. Happy Number
func isHappy(n int) bool {
	visited := make(map[int]bool)
	var getDigit func(n int) int
	getDigit = func(n int) int {
		res := 0
		for n > 0 {
			digit := n % 10
			res += digit * digit
			n /= 10
		}
		return res
	}

	for {
		if n == 1 {
			return true
		}
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = true
		n = getDigit(n)
	}

	return true
}

// 50. Pow(x, n)
// func MyPow(x float64, n int) float64 {

// }

// 66. Plus One
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] = digits[i] % 10
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
