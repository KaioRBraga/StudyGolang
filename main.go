package main

import "math"

func myAtoi(s string) int {
	n := len(s)

	if n == 0 {
		return 0
	}

	i := 0

	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	sinal := 1

	switch s[i] {
	case '-':
		sinal = -1
		i++
	case '+':
		i++
	}

	result := 0

	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		result = result*10 + digit

		if sinal*result > math.MaxInt32 {
			return math.MaxInt32
		}

		if sinal*result < math.MinInt32 {
			return math.MinInt32
		}

		i++
	}

	return sinal * result

}

func main() {
	s := "                          -4n2"
	result := myAtoi(s)
	println(result)
}
