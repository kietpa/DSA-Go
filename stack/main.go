package main

import "fmt"

type Stack struct {
	stack []int
}

func NewStack() Stack {
	return Stack{
		stack: []int{},
	}
}

func (s *Stack) AddToStack(item int) {
	s.stack = append(s.stack, item)
}

func (s *Stack) RemoveFromStack() (int, error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	// get last index for ans
	item := s.stack[len(s.stack)-1]

	// assign slice to exclude last index
	// since s[0:3] = s[0] s[1] s[2]
	s.stack = s.stack[0 : len(s.stack)-1]

	return item, nil
}

func main() {
	aa := NewStack()
	aa.AddToStack(1)
	aa.AddToStack(2)
	fmt.Println(aa)
	aa.RemoveFromStack()
	fmt.Println(aa)

	x := isValid("()")
	y := isValid("({})")
	z := isValid("([)]")

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z) // true, true, false
}

// VALID PARENTHESES

// Steps
// 1. check if open or closing bracket w/map
// 2. if open append to stack
// 3. if closed w/0 match in stack or empty stack false
// 4. if match w/open in stack (last index), remove open in stack & continue
// 5. if post loop len(stack) == 0 true
func isValid(s string) bool {
	stack := []rune{}
	braceMap := map[any]any{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		// 1
		// pair = open bracket if true
		pair, ok := braceMap[v]

		if !ok {
			// 2
			stack = append(stack, v)
		} else {
			// 3
			if len(stack) == 0 {
				return false
			}
			// 4
			if stack[len(stack)-1] != pair {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	// 5
	return len(stack) == 0
}
