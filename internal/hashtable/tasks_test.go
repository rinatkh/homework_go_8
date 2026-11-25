package hashtable

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
	"testing"

	"github.com/rinatkh/homework_go_8/internal/testutil"
)

func TestNoBuiltinMapInTable(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "table.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if _, ok := n.(*ast.MapType); ok {
			t.Error("own HashTable must not use builtin map")
		}
		return true
	})
}
func TestHashTable(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 9, 16, 31, 100, 1000} {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			for _, bad := range []bool{false, true} {
				var fn func(string) uint64
				if bad {
					fn = func(string) uint64 { return 0 }
				}
				table := New(4, fn)
				if table == nil {
					t.Fatal("constructor")
				}
				model := map[string]int{}
				for i := 0; i < n; i++ {
					key := fmt.Sprintf("key-%d", i)
					table.Set(key, i-4)
					model[key] = i - 4
				}
				if table.Len() != n || len(table.buckets) < 4 || float64(n)/float64(len(table.buckets)) > 0.75 {
					t.Fatal("len")
				}
				for k, v := range model {
					table.Set(k, v-1)
					model[k] = v - 1
				}
				for k, v := range model {
					if got, ok := table.Get(k); !ok || got != v {
						t.Fatal("lookup/update/collision/rehash")
					}
				}
				if table.Delete("absent") {
					t.Fatal("missing")
				}
				if _, ok := table.Get("absent"); ok {
					t.Fatal("missing lookup")
				}
				for k := range model {
					if !table.Delete(k) || table.Delete(k) {
						t.Fatal("delete")
					}
				}
				table.Set("", 0)
				if v, ok := table.Get(""); !ok || v != 0 || table.Len() != 1 {
					t.Fatal("empty key/zero/reuse")
				}
			}
		})
	}
}
func TestFrequency(t *testing.T) {
	for _, c := range testutil.Cases() {
		t.Run(c.Name, func(t *testing.T) {
			words := []string{}
			want := map[string]int{}
			for _, v := range c.Values {
				s := fmt.Sprint(v)
				words = append(words, s)
				want[s]++
			}
			before := slices.Clone(words)
			got := Frequencies(words)
			if len(got) != len(want) {
				t.Fatal("size")
			}
			for k, v := range want {
				if got[k] != v {
					t.Fatal("count")
				}
			}
			if !slices.Equal(words, before) {
				t.Fatal("input changed")
			}
		})
	}
}
func TestDuplicate(t *testing.T) {
	for _, c := range testutil.Cases() {
		t.Run(c.Name, func(t *testing.T) {
			before := slices.Clone(c.Values)
			want := false
			for i := range c.Values {
				for j := 0; j < i; j++ {
					if c.Values[i] == c.Values[j] {
						want = true
					}
				}
			}
			if ContainsDuplicate(c.Values) != want || !slices.Equal(before, c.Values) {
				t.Fatal("duplicate/unchanged")
			}
		})
	}
}

func TestSpecialKeys(t *testing.T) {
	for _, key := range []string{"", "a", "b", "aa", "ab", "ba", "zero", "0", "-1", "русский", "世界", "with space", "\x00", "last"} {
		t.Run(fmt.Sprintf("key=%q", key), func(t *testing.T) {
			table := New(4, func(string) uint64 { return 0 })
			table.Set("other", 9)
			table.Set(key, -7)
			table.Set(key, 0)
			if value, ok := table.Get(key); !ok || value != 0 || table.Len() != 2 {
				t.Fatal("special key/update")
			}
			if !table.Delete(key) || table.Delete(key) {
				t.Fatal("special key/delete")
			}
			if value, ok := table.Get("other"); !ok || value != 9 {
				t.Fatal("collision neighbor")
			}
		})
	}
}
