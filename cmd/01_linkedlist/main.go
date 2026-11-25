package main

import (
	"fmt"
	"os"

	"github.com/rinatkh/homework_go_8/internal/demo"
	"github.com/rinatkh/homework_go_8/internal/linkedlist"
)

func main() {
	head := &linkedlist.Node{Value: 10, Next: &linkedlist.Node{Value: 20, Next: &linkedlist.Node{Value: 30}}}
	reversed := linkedlist.Reverse(head)
	// Только сериализация для вывода, не решение Reverse. Лимит защищает
	// демонстрацию от бесконечного обхода ошибочно зацикленного результата.
	values := []int{}
	p := reversed
	for len(values) < 8 && p != nil {
		values = append(values, p.Value)
		p = p.Next
	}
	if p != nil {
		fmt.Println("Вывод цепочки ограничен 8 узлами; проверь ссылки.")
	}
	demo.Print(os.Stdout, "Reverse", "10 → 20 → 30 → nil", values, []int{30, 20, 10})
	first, second := &linkedlist.Node{Value: 1}, &linkedlist.Node{Value: 2}
	first.Next = second
	second.Next = first
	demo.Print(os.Stdout, "HasCycle", "1 → 2 → 1 (цикл)", linkedlist.HasCycle(first), true)
}
