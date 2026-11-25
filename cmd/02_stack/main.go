package main

import (
	"os"

	"github.com/rinatkh/homework_go_8/internal/demo"
	"github.com/rinatkh/homework_go_8/internal/stack"
)

func main() {
	demo.Print(os.Stdout, "ValidParentheses", "([]{})", stack.ValidParentheses("([]{})"), true)
	var s stack.MinStack
	s.Push(3)
	s.Push(1)
	s.Push(1)
	value, ok := s.GetMin()
	type result struct {
		Value int
		OK    bool
	}
	demo.Print(os.Stdout, "MinStack.GetMin", "Push(3), Push(1), Push(1)", result{value, ok}, result{1, true})
	s.Pop()
	value, ok = s.GetMin()
	demo.Print(os.Stdout, "Повторный минимум", "затем Pop()", result{value, ok}, result{1, true})
}
