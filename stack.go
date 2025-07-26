package main

import (
	"strconv"
	"strings"
)

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

// 155. Min Stack
type MinStack struct {
	Stack    []int
	MinStack []int
}

func MinStackConstructor() MinStack {
	return MinStack{
		Stack:    []int{},
		MinStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.MinStack) == 0 || val <= this.MinStack[len(this.MinStack)-1] {
		this.MinStack = append(this.MinStack, val)
	} else {
		this.MinStack = append(this.MinStack, this.MinStack[len(this.MinStack)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) > 0 {
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	if len(this.MinStack) > 0 {
		this.MinStack = this.MinStack[:len(this.MinStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.Stack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// 22. Generate Parentheses
func GenerateParenthesis(n int) []string {
	// BRUTE FORCE
	// 	res := make([]string, 0)

	//    var valid func(string) bool
	//    valid = func(s string) bool {
	//        open := 0
	//        for _, c := range s {
	//            if c == '(' {
	//                open++
	//            } else {
	//                open--
	//            }
	//            if open < 0 {
	//                return false
	//            }
	//        }
	//        return open == 0
	//    }

	//    var dfs func(string)
	//    dfs = func(s string) {
	//        if len(s) == n*2 {
	//            if valid(s) {
	//                res = append(res, s)
	//            }
	//            return
	//        }

	//        dfs(s + "(")
	//        dfs(s + ")")
	//    }

	//    dfs("")
	//    return res
	stack := []string{}
	res := []string{}
	var backtrack func(openN, closedN int)
	backtrack = func(openN, closedN int) {
		if openN == closedN && openN == n {
			p := strings.Join(stack, "")
			res = append(res, p)
		}
		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}

		if closedN < openN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return res
}
