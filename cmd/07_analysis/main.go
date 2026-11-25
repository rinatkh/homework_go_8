package main

import (
	"os"

	"github.com/rinatkh/homework_go_8/internal/analysis"
	"github.com/rinatkh/homework_go_8/internal/demo"
)

func main() {
	demo.Print(os.Stdout, "Analyze", "list-insert-after, knownHandle=true", analysis.Analyze("list-insert-after", true), analysis.Cost{Time: "O(1)", Extra: "O(1)", Total: "O(n)"})
	demo.Print(os.Stdout, "SelectStructure", "fifo-fixed-reject", analysis.SelectStructure("fifo-fixed-reject"), "ring")
}
