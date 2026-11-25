// Package linkedlist — самостоятельные задачи, решения остаются TODO.
package linkedlist

// Node используется во всех пяти задачах. На входе Reverse/Merge/Remove — ациклические цепочки.
type Node struct {
	Value int
	Next  *Node
}

// Reverse перестраивает связи существующих узлов, O(n)/O(1).
func Reverse(head *Node) *Node { // TODO: 01.1, LeetCode 206.
	return nil
}

// MergeSorted объединяет две независимые sorted цепочки, переиспользуя узлы.
// Равные узлы first идут раньше second. O(n+m)/O(1), без нового массива значений.
func MergeSorted(first, second *Node) *Node { // TODO: 01.2, LeetCode 21.
	return nil
}
func HasCycle(head *Node) bool { // TODO: 01.3, LeetCode 141, O(n)/O(1).
	return false
}

// RemoveValue удаляет все совпадения и очищает Next у отсоединённых узлов.
func RemoveValue(head *Node, value int) *Node { // TODO: 01.4, O(n)/O(1).
	return nil
}

// InsertAfter не ищет узел: nil отклоняется; иначе создаёт ровно один новый узел.
func InsertAfter(at *Node, value int) (*Node, bool) { // TODO: 01.5, O(1).
	return nil, false
}
