package queue

// TwoStackQueue: каждый элемент переносится из incoming в outgoing не более раза.
// Нельзя заменять задачу одним срезом с copy. Zero value готов к работе.
type TwoStackQueue struct{ incoming, outgoing []int }

func (q *TwoStackQueue) Push(value int) { // TODO: 03.1, LeetCode 232.
}
func (q *TwoStackQueue) Pop() (int, bool)  { return 0, false } // TODO: амортизированно O(1).
func (q *TwoStackQueue) Peek() (int, bool) { return 0, false } // TODO: не удалять front.
func (q *TwoStackQueue) Len() int          { return 0 }        // TODO.
// RoundRobin возвращает steps имён: каждое взятое имя уходит в конец очереди.
// При steps<=0 либо пустом names результат пуст. names не меняется.
func RoundRobin(names []string, steps int) []string { // TODO: 03.2, O(n+steps).
	return nil
}

func (q *TwoStackQueue) Empty() bool { return false } // TODO: LeetCode 232.
