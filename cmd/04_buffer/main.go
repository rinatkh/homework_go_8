package main

import (
	"fmt"
	"os"

	"github.com/rinatkh/homework_go_8/internal/buffer"
	"github.com/rinatkh/homework_go_8/internal/demo"
)

func main() {
	var b buffer.ByteBuffer
	n, err := b.Write([]byte("Go API"))
	type writeResult struct {
		N     int
		Error string
	}
	demo.Print(os.Stdout, "ByteBuffer.Write", "Go API", writeResult{n, fmt.Sprint(err)}, writeResult{6, "<nil>"})
	demo.Print(os.Stdout, "ByteBuffer.Bytes", "после Write", string(b.Bytes()), "Go API")
	data := make([]byte, 3)
	n, err = b.Read(data)
	text := "некорректная длина Read"
	if n >= 0 && n <= len(data) {
		text = string(data[:n])
	}
	type readResult struct {
		N           int
		Text, Error string
	}
	demo.Print(os.Stdout, "ByteBuffer.Read", "буфер чтения 3 байта", readResult{n, text, fmt.Sprint(err)}, readResult{3, "Go ", "<nil>"})
}
