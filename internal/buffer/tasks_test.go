package buffer

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestByteBuffer(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 9, 16, 31, 100, 1000} {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			b := new(ByteBuffer)
			input := bytes.Repeat([]byte("x"), n)
			got, err := b.Write(input)
			if got != n || err != nil || b.Len() != n {
				t.Fatal("write")
			}
			if n > 0 {
				input[0] = 'Y'
			}
			if !bytes.Equal(b.Bytes(), bytes.Repeat([]byte("x"), n)) {
				t.Fatal("write alias")
			}
			copy := b.Bytes()
			if len(copy) > 0 {
				copy[0] = 'Z'
				if b.Bytes()[0] != 'x' {
					t.Fatal("Bytes alias")
				}
			}
			if read, e := b.Read(nil); read != 0 || e != nil {
				t.Fatal("Read(nil)")
			}
			out := make([]byte, n+1)
			read, e := b.Read(out)
			if n == 0 {
				if read != 0 || e != io.EOF {
					t.Fatal("EOF")
				}
			} else if read != n || e != nil || !bytes.Equal(out[:read], bytes.Repeat([]byte("x"), n)) {
				t.Fatal("Read")
			}
			b.Write([]byte("abc"))
			oldCap := cap(b.data)
			backing := b.data
			b.Reset()
			if b.Len() != 0 || cap(b.data) != oldCap {
				t.Fatal("Reset")
			}
			for _, v := range backing {
				if v != 0 {
					t.Fatal("clear")
				}
			}
		})
	}
}
func TestAppendBounded(t *testing.T) {
	for _, c := range []struct {
		name       string
		dst, chunk string
		limit      int
		ok         bool
	}{{"nil_empty", "", "", 0, true}, {"one", "", "a", 1, true}, {"negative_limit", "", "", -1, false}, {"full_reject", "abc", "d", 3, false}, {"exact", "ab", "cd", 4, true}, {"spare", "ab", "c", 9, true}, {"chunk_too_big", "a", "bcdef", 3, false}, {"existing_over_limit", "abcd", "", 3, false}, {"empty_chunk", "ab", "", 2, true}, {"zero_limit", "", "a", 0, false}, {"unicode_bytes", "я", "!", 3, true}, {"unicode_reject", "я", "!", 2, false}, {"zeros", "\x00", "\x00", 2, true}} {
		t.Run(c.name, func(t *testing.T) {
			dst, chunk := []byte(c.dst), []byte(c.chunk)
			before := bytes.Clone(dst)
			got, ok := AppendBounded(dst, chunk, c.limit)
			want := c.dst
			if c.ok {
				want += c.chunk
			}
			if ok != c.ok || string(got) != want || string(chunk) != c.chunk || (!ok && !bytes.Equal(dst, before)) {
				t.Fatal("bounded atomic append")
			}
		})
	}
}
