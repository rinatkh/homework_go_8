# 05. ringbuffer — 2 задач

Функциональные тесты находятся рядом с TODO в internal/. Решения здесь не приведены. nil и пустой результат эквивалентны по элементам, если явно не требуется handle.

## 05.1 CircularQueue

```go
func New(capacity int) (*CircularQueue,error); Enqueue(int) bool; Dequeue() (int,bool); Front()/Rear() (int,bool); Len() int; IsEmpty()/IsFull() bool
```

**Контракт:** Фиксированная FIFO, positive capacity; full отклоняет новую запись, существующие не теряются.

**Пример:** `cap=2: enqueue1,2; enqueue3=false; dequeue=1; enqueue3; front=2,rear=3`.

**Ограничения:** Capacity≤10000; capacity≤0 -> nil,ErrCapacity. В LC DeQueue возвращает bool и Front/Rear -1; здесь Dequeue возвращает значение и ok, чтобы хранить отрицательные.

**Подсказка:** head==tail бывает у full и empty: какая дополнительная информация нужна?

**Ожидаемая сложность:** Операции O(1), New O(cap), память O(cap).

**Тесты:** 13 capacity-сценариев, каждый до 1000 операций, wrap, full/empty/front/rear.

[Оригинал LeetCode](https://leetcode.com/problems/design-circular-queue/). Условия выше — учебная адаптация с явно указанными изменениями API.

## 05.2 RecentValues

```go
func RecentValues(values []int, k int) []int
```

**Контракт:** Вернуть последние k значений в порядке поступления, используя кольцо с вытеснением старого. Результат — копия.

**Пример:** `[1,2,3,4,5],3 -> [3,4,5]`.

**Ограничения:** n,k≤10000, k≤0 -> пусто, values не менять; не ограничиваться одним viewslice входа.

**Подсказка:** При full политика отличается от classroom: какое значение исчезает?

**Ожидаемая сложность:** O(n) времени, O(k) дополнительной памяти и результата.

**Тесты:** 15 наборов × 6 k, проверка порядка, unchanged и aliasing.
