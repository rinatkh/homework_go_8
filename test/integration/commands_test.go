package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// Тестирует запуск и контракт вывода, не закрепляя неверные ответы заглушек.
func TestCommands(t *testing.T) {
	for _, test := range []struct{ command, example string }{
		{"01_linkedlist", "Reverse"}, {"02_stack", "ValidParentheses"}, {"03_queue", "TwoStackQueue.Pop"},
		{"04_buffer", "ByteBuffer.Write"}, {"05_ringbuffer", "CircularQueue.Enqueue"},
		{"06_hashtable", "HashTable.Get"}, {"07_analysis", "Analyze"}, {"08_benchmark", "Объявление найдено: true"},
	} {
		t.Run(test.command, func(t *testing.T) {
			command := exec.Command("go", "run", "./cmd/"+test.command)
			command.Dir = "../.."
			output, err := command.CombinedOutput()
			if err != nil {
				t.Fatalf("go run: %v\n%s", err, output)
			}
			text := string(output)
			if !strings.Contains(text, test.example) || strings.Contains(text, "SUCCESS") {
				t.Fatal(text)
			}
			if test.command == "08_benchmark" {
				if !strings.Contains(text, "Измерения сейчас не выполнялись.") || strings.Contains(text, "ns/op") {
					t.Fatal("unexpected measurement:", text)
				}
				return
			}
			cases := strings.Count(text, "Алгоритм:")
			if cases == 0 || strings.Count(text, "Получено:") != cases || strings.Count(text, "Ожидается:") != cases || strings.Count(text, "Результаты отличаются")+strings.Count(text, "Пример совпал.") != cases {
				t.Fatal(text)
			}
		})
	}
}
