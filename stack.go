package main

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
