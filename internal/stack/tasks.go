package stack

// ValidParentheses принимает только ()[]{}; другой символ — false. Пустая строка корректна.
func ValidParentheses(text string) bool { // TODO: 02.1, LeetCode 20.
	return false
}

// MinStack хранит дубликаты минимумов. Zero value готов к работе.
// Все операции O(1) амортизированно; Pop/Top/GetMin на пустом возвращают (0,false).
type MinStack struct{ values, minimums []int }

func (s *MinStack) Push(value int) { // TODO: 02.2, LeetCode 155.
}
func (s *MinStack) Pop() (int, bool)    { return 0, false } // TODO: сохранить повторный минимум.
func (s *MinStack) Top() (int, bool)    { return 0, false } // TODO.
func (s *MinStack) GetMin() (int, bool) { return 0, false } // TODO.
func (s *MinStack) Len() int            { return 0 }        // TODO.
