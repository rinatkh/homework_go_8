package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Лаборатория benchmark: диагностика объявления, без измерения скорости.")
	// В пакете есть только _test.go. Discovery компилирует его, но не запускает
	// benchmark и не оценивает готовность TODO по наличию имени или комментария.
	command := exec.Command("go", "test", "-list", "^BenchmarkStructures$", "./internal/benchmark")
	output, err := command.CombinedOutput()
	fmt.Printf("Фактический вывод Go:\n%s", output)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка компиляции/discovery:", err)
		os.Exit(1)
	}
	found := false
	for _, line := range strings.Split(string(output), "\n") {
		if line == "BenchmarkStructures" {
			found = true
		}
	}
	fmt.Printf("Объявление найдено: %t; ожидается: true.\n", found)
	fmt.Println("Это не проверка реализации. Измерения сейчас не выполнялись.")
	fmt.Println("Проверить TODO одним вызовом: make test-benchmark")
	fmt.Println("После реализации структур и лаборатории: make bench BENCHTIME=100ms BENCHCOUNT=3")
	if !found {
		os.Exit(1)
	}
}
