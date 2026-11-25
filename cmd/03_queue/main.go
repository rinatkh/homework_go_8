package main

import (
	"os"

	"github.com/rinatkh/homework_go_8/internal/demo"
	"github.com/rinatkh/homework_go_8/internal/queue"
)

func main() {
	var q queue.TwoStackQueue
	q.Push(10)
	q.Push(20)
	q.Push(30)
	value, ok := q.Pop()
	type result struct {
		Value int
		OK    bool
	}
	demo.Print(os.Stdout, "TwoStackQueue.Pop", "Push(10), Push(20), Push(30)", result{value, ok}, result{10, true})
	demo.Print(os.Stdout, "RoundRobin", "names=[A B], steps=5", queue.RoundRobin([]string{"A", "B"}, 5), []string{"A", "B", "A", "B", "A"})
}
