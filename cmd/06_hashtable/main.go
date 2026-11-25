package main

import (
	"os"

	"github.com/rinatkh/homework_go_8/internal/demo"
	"github.com/rinatkh/homework_go_8/internal/hashtable"
)

func main() {
	table := hashtable.New(4, nil)
	if table == nil {
		demo.Print(os.Stdout, "HashTable.New", "capacity=4", false, true)
		return
	}
	table.Set("alice", 7)
	value, ok := table.Get("alice")
	type result struct {
		Value int
		OK    bool
	}
	demo.Print(os.Stdout, "HashTable.Get", "Set(alice,7), Get(alice)", result{value, ok}, result{7, true})
	demo.Print(os.Stdout, "Frequencies", []string{"go", "api", "go"}, hashtable.Frequencies([]string{"go", "api", "go"}), map[string]int{"go": 2, "api": 1})
	demo.Print(os.Stdout, "ContainsDuplicate", []int{4, 2, 4}, hashtable.ContainsDuplicate([]int{4, 2, 4}), true)
}
