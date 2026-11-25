package main

import (
	"fmt"
	"os"

	"github.com/rinatkh/homework_go_8/internal/demo"
	"github.com/rinatkh/homework_go_8/internal/ringbuffer"
)

func main() {
	q, err := ringbuffer.New(3)
	if q == nil || err != nil {
		demo.Print(os.Stdout, "CircularQueue.New", "capacity=3", fmt.Sprintf("queue=%v, err=%v", q, err), "не nil, err=<nil>")
		return
	}
	accepted := []bool{q.Enqueue(10), q.Enqueue(20), q.Enqueue(30), q.Enqueue(40)}
	demo.Print(os.Stdout, "CircularQueue.Enqueue", "capacity=3, значения 10/20/30/40", accepted, []bool{true, true, true, false})
	value, ok := q.Dequeue()
	type result struct {
		Value int
		OK    bool
	}
	demo.Print(os.Stdout, "CircularQueue.Dequeue", "после записи", result{value, ok}, result{10, true})
	demo.Print(os.Stdout, "RecentValues", "values=[1 2 3 4 5], k=3", ringbuffer.RecentValues([]int{1, 2, 3, 4, 5}, 3), []int{3, 4, 5})
}
