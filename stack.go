package main

import "strconv"

// 20. Valid Parentheses
func IsValid(s string) bool {

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) < 1 {
				return false
			}
			top := stack[len(stack)-1]
			if (top == '(' && s[i] == ')') || (top == '[' && s[i] == ']') || (top == '{' && s[i] == '}') {
				stack = stack[:len(stack)-1]
				continue
			}
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 739. Daily Temperatures
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	monotonicStack := []int{}
	for i, temp := range temperatures {
		for len(monotonicStack) > 0 && temperatures[monotonicStack[len(monotonicStack)-1]] < temp {
			res[monotonicStack[len(monotonicStack)-1]] = i - monotonicStack[len(monotonicStack)-1]
			monotonicStack = monotonicStack[:len(monotonicStack)-1]
		}
		monotonicStack = append(monotonicStack, i)
	}
	return res
}

// 150. Evaluate Reverse Polish Notation
func evalRPN(tokens []string) int {
	// stack := []int{}
	// for _, token := range tokens {
	// 	if num, err := strconv.Atoi(token); err == nil {
	// 		stack = append(stack, num)
	// 	} else {
	// 		b := stack[len(stack)-1]
	// 		a := stack[len(stack)-2]
	// 		stack = stack[:len(stack)-2]
	// 		if token == "+" {
	// 			stack = append(stack, a+b)
	// 		} else if token == "-" {
	// 			stack = append(stack, a-b)
	// 		} else if token == "*" {
	// 			stack = append(stack, a*b)
	// 		} else {
	// 			stack = append(stack, a/b)
	// 		}
	// 	}
	// }
	// return stack[0]
	stack := []int{}
	for _, token := range tokens {
		if token == "+" {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		} else if token == "-" {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		} else if token == "*" {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		} else if token == "/" {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		} else {
			if num, err := strconv.Atoi(token); err == nil {
				stack = append(stack, num)
			}
		}
	}
	return stack[0]
}
