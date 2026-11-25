package analysis

import (
	"fmt"
	"testing"
)

func TestAnalyze(t *testing.T) {
	for _, c := range []struct {
		op    string
		known bool
		want  Cost
	}{
		{"list-find", false, Cost{"O(n)", "O(1)", "O(n)"}},
		{"list-insert-after", true, Cost{"O(1)", "O(1)", "O(n)"}},
		{"list-insert-after", false, Cost{"O(n)", "O(1)", "O(n)"}},
		{"list-reverse", true, Cost{"O(n)", "O(1)", "O(n)"}},
		{"slice-index", true, Cost{"O(1)", "O(1)", "O(cap)"}},
		{"slice-prepend", false, Cost{"O(n)", "O(n)", "O(cap)"}},
		{"slice-stack-push", true, Cost{"amortized O(1)", "O(n)", "O(cap)"}},
		{"linked-queue-dequeue", true, Cost{"O(1)", "O(1)", "O(n)"}},
		{"ring-enqueue", true, Cost{"O(1)", "O(1)", "O(cap)"}},
		{"hash-get", true, Cost{"expected O(1)", "O(1)", "O(n+cap)"}},
		{"hash-set-growth", true, Cost{"O(n+cap)", "O(cap)", "O(n+cap)"}},
		{"unknown", false, Cost{"unknown", "unknown", "unknown"}},
	} {
		t.Run(fmt.Sprintf("%s/known=%t", c.op, c.known), func(t *testing.T) {
			if got := Analyze(c.op, c.known); got != c.want {
				t.Fatalf("got %+v want %+v", got, c.want)
			}
		})
	}
}
func TestSelect(t *testing.T) {
	for _, c := range []struct{ scenario, want string }{
		{"random-index", "slice"},
		{"lifo-undo", "stack"},
		{"fifo-unbounded", "linked-queue"},
		{"fifo-fixed-reject", "ring"},
		{"append-bytes", "buffer"},
		{"lookup-key", "map"},
		{"remove-known-node-both-directions", "doubly-list"},
		{"reverse-chain", "singly-list"},
		{"retain-last-k", "ring"},
		{"dense-indexes", "slice"},
		{"pending-lifo", "stack"},
		{"unknown", "unknown"},
	} {
		t.Run(c.scenario, func(t *testing.T) {
			if SelectStructure(c.scenario) != c.want {
				t.Fatal("structure choice")
			}
		})
	}
}
