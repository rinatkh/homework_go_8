package buffer

import "io"

// ByteBuffer реализует io.Reader/io.Writer: Read(nil) -> (0,nil), empty -> io.EOF.
// Write копирует вход, Bytes возвращает независимую копию, Reset сохраняет capacity.
// Стандартный bytes.Buffer можно использовать как oracle в tests, не как решение.
type ByteBuffer struct {
	data   []byte
	offset int
}

var _ io.Reader = (*ByteBuffer)(nil)
var _ io.Writer = (*ByteBuffer)(nil)

func (b *ByteBuffer) Write(p []byte) (int, error) { return 0, nil } // TODO: 04.1.
func (b *ByteBuffer) Read(p []byte) (int, error)  { return 0, nil } // TODO.
func (b *ByteBuffer) Bytes() []byte               { return nil }    // TODO.
func (b *ByteBuffer) Len() int                    { return 0 }      // TODO.
func (b *ByteBuffer) Reset()                      {}                // TODO.
// AppendBounded: limit ограничивает итоговую длину, limit<0 — отказ.
// При отказе вернуть dst без изменений и false; частичная запись запрещена.
// При успехе append; вход chunk не менять; alias dst допустим, как у append.
func AppendBounded(dst, chunk []byte, limit int) ([]byte, bool) { // TODO: 04.2.
	return dst, false
}
