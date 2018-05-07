package main

import "fmt"

func main() {
	a := "{[](){}}"
	fmt.Println(checkifbracketmatches(a))
}

func checkifbracketmatches(a string) bool {
	bytesss := []byte(a)
	stack := make([]string, 0)
	for _, c := range bytesss {
		if string(c) == "}" {
			if len(stack) > 0 {
				if stack[len(stack)-1] == "{" {
					stack = stack[:len(stack)-1]
					continue
				}
			}
		}
		if string(c) == ")" {
			if len(stack) > 0 {
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
					continue
				}
			}
		}
		if string(c) == "]" {
			if len(stack) > 0 {
				if stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
					continue
				}
			}
		}
		stack = append(stack, string(c))
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
