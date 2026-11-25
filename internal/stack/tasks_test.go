package stack

import (
	"fmt"
	"slices"
	"testing"

	"github.com/rinatkh/homework_go_8/internal/testutil"
)

func TestParentheses(t *testing.T) {
	for _, c := range []struct {
		name, text string
		ok         bool
	}{{"empty", "", true}, {"round", "()", true}, {"square", "[]", true}, {"curly", "{}", true}, {"nested", "({[]})", true}, {"adjacent", "()[]{}", true}, {"crossed", "([)]", false}, {"extra_close", "())", false}, {"unfinished", "(()", false}, {"only_open", "[", false}, {"only_close", "}", false}, {"letters", "go", false}, {"unicode", "я", false}, {"deep", "(((())))", true}} {
		t.Run(c.name, func(t *testing.T) {
			if ValidParentheses(c.text) != c.ok {
				t.Fatal("parentheses")
			}
		})
	}
}
func TestMinStack(t *testing.T) {
	for _, c := range testutil.Cases() {
		t.Run(c.Name, func(t *testing.T) {
			s := new(MinStack)
			if v, ok := s.Pop(); ok || v != 0 {
				t.Fatal("empty pop")
			}
			if _, ok := s.Top(); ok {
				t.Fatal("empty top")
			}
			if _, ok := s.GetMin(); ok {
				t.Fatal("empty min")
			}
			model := []int{}
			for _, v := range c.Values {
				s.Push(v)
				model = append(model, v)
				if got, ok := s.GetMin(); !ok || got != slices.Min(model) {
					t.Fatal("min")
				}
				if top, ok := s.Top(); !ok || top != v || s.Len() != len(model) {
					t.Fatal("top/len")
				}
			}
			for len(model) > 0 {
				want := model[len(model)-1]
				got, ok := s.Pop()
				if !ok || got != want {
					t.Fatal("LIFO")
				}
				model = model[:len(model)-1]
				m, valid := s.GetMin()
				if valid != (len(model) > 0) || (valid && m != slices.Min(model)) {
					t.Fatal("min after pop")
				}
			}
			s.Push(-7)
			if m, ok := s.GetMin(); !ok || m != -7 {
				t.Fatal(fmt.Sprint("reuse", m, ok))
			}
		})
	}
}
