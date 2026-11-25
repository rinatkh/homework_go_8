package queue

import (
	"fmt"
	"slices"
	"testing"

	"github.com/rinatkh/homework_go_8/internal/testutil"
)

func TestTwoStacks(t *testing.T) {
	for _, c := range testutil.Cases() {
		t.Run(c.Name, func(t *testing.T) {
			q := new(TwoStackQueue)
			if !q.Empty() {
				t.Fatal("initial empty")
			}
			if _, ok := q.Pop(); ok {
				t.Fatal("empty")
			}
			model := []int{}
			for _, v := range c.Values {
				q.Push(v)
				model = append(model, v)
			}
			for i := 0; i < len(c.Values)+3; i++ {
				if i%2 == 0 {
					q.Push(-91 - i)
					model = append(model, -91-i)
				}
				got, ok := q.Peek()
				if ok != (len(model) > 0) || (ok && got != model[0]) {
					t.Fatal("front")
				}
				v, valid := q.Pop()
				if valid != ok || (ok && v != got) {
					t.Fatal("pop")
				}
				if ok {
					model = model[1:]
				}
				if q.Len() != len(model) || q.Empty() != (len(model) == 0) {
					t.Fatal("len")
				}
			}
		})
	}
}
func TestRoundRobin(t *testing.T) {
	for _, c := range []struct {
		name  string
		names []string
		steps int
	}{{"nil", nil, 4}, {"empty", []string{}, 4}, {"zero", []string{"a"}, 0}, {"negative", []string{"a"}, -1}, {"one", []string{"a"}, 1}, {"repeat_one", []string{"a"}, 5}, {"pair", []string{"a", "b"}, 2}, {"partial", []string{"a", "b", "c"}, 2}, {"wrap", []string{"a", "b", "c"}, 8}, {"duplicate", []string{"a", "a", "b"}, 7}, {"unicode", []string{"А", "Б"}, 5}, {"many", []string{"a", "b"}, 1000}} {
		t.Run(c.name, func(t *testing.T) {
			before := slices.Clone(c.names)
			var want []string
			if len(c.names) > 0 {
				for i := 0; i < c.steps; i++ {
					want = append(want, c.names[i%len(c.names)])
				}
			}
			got := RoundRobin(c.names, c.steps)
			if !slices.Equal(got, want) || !slices.Equal(before, c.names) {
				t.Fatal(fmt.Sprint("round robin: ", got))
			}
		})
	}
}
