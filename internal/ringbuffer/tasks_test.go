package ringbuffer

import (
	"fmt"
	"slices"
	"testing"

	"github.com/rinatkh/homework_go_8/internal/testutil"
)

func TestCircularQueue(t *testing.T) {
	for _, capacity := range []int{-1, 0, 1, 2, 3, 4, 5, 7, 8, 9, 16, 31, 64} {
		t.Run(fmt.Sprint(capacity), func(t *testing.T) {
			q, err := New(capacity)
			if capacity <= 0 {
				if err != ErrCapacity || q != nil {
					t.Fatal("invalid capacity")
				}
				return
			}
			if err != nil || q == nil {
				t.Fatal("constructor")
			}
			if _, ok := q.Front(); ok {
				t.Fatal("empty front")
			}
			if _, ok := q.Dequeue(); ok {
				t.Fatal("empty dequeue")
			}
			if !q.IsEmpty() {
				t.Fatal("initial empty")
			}
			model := []int{}
			for step := 0; step < 1000; step++ {
				if step%3 != 0 {
					value := step%9 - 4
					ok := q.Enqueue(value)
					if ok != (len(model) < capacity) {
						t.Fatal("full")
					}
					if ok {
						model = append(model, value)
					}
				} else {
					v, ok := q.Dequeue()
					if ok != (len(model) > 0) || (ok && v != model[0]) {
						t.Fatal("FIFO")
					}
					if ok {
						model = model[1:]
					}
				}
				if q.Len() != len(model) || q.IsFull() != (len(model) == capacity) || q.IsEmpty() != (len(model) == 0) {
					t.Fatal("len/full")
				}
				last, present := q.Rear()
				if present != (len(model) > 0) || (present && last != model[len(model)-1]) {
					t.Fatal("rear")
				}
				v, ok := q.Front()
				if ok != (len(model) > 0) || (ok && v != model[0]) {
					t.Fatal("front")
				}
				if q.head < 0 || q.head >= capacity || q.tail < 0 || q.tail >= capacity {
					t.Fatal("index")
				}
			}
		})
	}
}
func TestRecentValues(t *testing.T) {
	for _, c := range testutil.Cases() {
		for _, k := range []int{-1, 0, 1, 2, 3, 20} {
			t.Run(fmt.Sprintf("%s/k=%d", c.Name, k), func(t *testing.T) {
				before := slices.Clone(c.Values)
				var want []int
				if k > 0 {
					start := max(0, len(c.Values)-k)
					want = c.Values[start:]
				}
				got := RecentValues(c.Values, k)
				if !slices.Equal(got, want) || !slices.Equal(c.Values, before) {
					t.Fatal("recent/order/unchanged")
				}
				if len(got) > 0 {
					got[0] = 1234567
					if !slices.Equal(c.Values, before) {
						t.Fatal("alias")
					}
				}
			})
		}
	}
}
