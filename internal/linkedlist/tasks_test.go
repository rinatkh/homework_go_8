package linkedlist

import (
	"fmt"
	"slices"
	"testing"

	"github.com/rinatkh/homework_go_8/internal/testutil"
)

func chain(v []int) (*Node, []*Node) {
	nodes := make([]*Node, len(v))
	for i, x := range v {
		nodes[i] = &Node{Value: x}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	if len(nodes) == 0 {
		return nil, nodes
	}
	return nodes[0], nodes
}
func collect(t *testing.T, head *Node, limit int) []*Node {
	t.Helper()
	var out []*Node
	for p := head; p != nil; p = p.Next {
		if len(out) > limit {
			t.Fatal("unexpected cycle")
		}
		out = append(out, p)
	}
	return out
}
func values(nodes []*Node) []int {
	var out []int
	for _, n := range nodes {
		out = append(out, n.Value)
	}
	return out
}
func TestListTasks(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			t.Run("Reverse", func(t *testing.T) {
				head, nodes := chain(tc.Values)
				got := collect(t, Reverse(head), len(nodes))
				slices.Reverse(nodes)
				if !slices.Equal(got, nodes) {
					t.Fatal("reverse must reuse nodes in reverse order")
				}
			})
			t.Run("RemoveValue", func(t *testing.T) {
				for _, target := range append(slices.Clone(tc.Values), 1234567) {
					head, nodes := chain(tc.Values)
					var want []*Node
					for _, n := range nodes {
						if n.Value != target {
							want = append(want, n)
						}
					}
					got := collect(t, RemoveValue(head, target), len(nodes))
					if !slices.Equal(got, want) {
						t.Fatal("remove all/identity")
					}
					for _, n := range nodes {
						if n.Value == target && n.Next != nil {
							t.Fatal("deleted link retained")
						}
					}
				}
			})
			t.Run("InsertAfter", func(t *testing.T) {
				head, nodes := chain(tc.Values)
				_ = head
				var at *Node
				if len(nodes) > 0 {
					at = nodes[len(nodes)/2]
				}
				var old *Node
				if at != nil {
					old = at.Next
				}
				n, ok := InsertAfter(at, -91)
				if ok != (at != nil) {
					t.Fatal("nil contract")
				}
				if ok && (n == nil || n == at || n.Value != -91 || at.Next != n || n.Next != old) {
					t.Fatal("insert linkage")
				}
			})
		})
	}
}
func TestMergeSorted(t *testing.T) {
	for _, tc := range testutil.Cases() {
		t.Run(tc.Name, func(t *testing.T) {
			a := slices.Clone(tc.Values)
			slices.Sort(a)
			b := []int{-2, 0, 4}
			first, aa := chain(a)
			second, bb := chain(b)
			expected := append(slices.Clone(aa), bb...)
			slices.SortStableFunc(expected, func(a, b *Node) int {
				if a.Value < b.Value {
					return -1
				}
				if a.Value > b.Value {
					return 1
				}
				return 0
			})
			got := collect(t, MergeSorted(first, second), len(expected))
			if !slices.Equal(got, expected) {
				t.Fatal("stable merge must reuse exact nodes")
			}
		})
	}
	for _, pair := range []struct {
		name          string
		first, second []int
	}{{"both_nil", nil, nil}, {"right_nil", []int{1, 2}, nil}, {"left_nil", nil, []int{1, 2}}} {
		t.Run(pair.name, func(t *testing.T) {
			a, aa := chain(pair.first)
			b, bb := chain(pair.second)
			got := collect(t, MergeSorted(a, b), len(aa)+len(bb))
			want := append(aa, bb...)
			if !slices.Equal(got, want) {
				t.Fatal("empty side")
			}
		})
	}
}
func TestCycle(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 9, 16, 31, 100} {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			for entry := -1; entry < n; entry++ {
				head, nodes := chain(make([]int, n))
				if n > 0 && entry >= 0 {
					nodes[n-1].Next = nodes[entry]
				}
				before := make([]*Node, n)
				for i, p := range nodes {
					before[i] = p.Next
				}
				if HasCycle(head) != (entry >= 0) {
					t.Fatalf("entry=%d", entry)
				}
				for i, p := range nodes {
					if p.Next != before[i] {
						t.Fatal("input changed")
					}
				}
			}
		})
	}
}
