package ringbuffer

import "errors"

var ErrCapacity = errors.New("capacity must be positive")

// CircularQueue использует ровно capacity слотов, не увеличивает массив.
// Для full отклонить запись, empty — (0,false), negative values допустимы.
type CircularQueue struct {
	slots            []int
	head, tail, size int
}

func New(capacity int) (*CircularQueue, error) { // TODO: 05.1, LeetCode 622 (адаптация API).
	return &CircularQueue{}, nil
}
func (q *CircularQueue) Enqueue(value int) bool { return false }    // TODO.
func (q *CircularQueue) Dequeue() (int, bool)   { return 0, false } // TODO, очистить слот.
func (q *CircularQueue) Front() (int, bool)     { return 0, false } // TODO.
func (q *CircularQueue) Len() int               { return 0 }        // TODO.
func (q *CircularQueue) IsFull() bool           { return false }    // TODO.
// RecentValues хранит последние k значений: при full вытесняет СТАРОЕ.
// Это другая политика, чем classroom Ring. Вернуть копию в порядке поступления.
func RecentValues(values []int, k int) []int { // TODO: 05.2, O(n), O(k).
	return nil
}

func (q *CircularQueue) Rear() (int, bool) { return 0, false } // TODO: последний элемент.
func (q *CircularQueue) IsEmpty() bool     { return false }    // TODO.
