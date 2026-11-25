# 07. Анализ и выбор — 2 задачи

## 07.1 Analyze

`func Analyze(operation string, knownHandle bool) Cost` возвращает Time/Extra/Total — канонические метки из tasks_test.go. Здесь Extra — **худшая** дополнительная память одного вызова, Total — суммарная структура. Для slice append Time амортизированное, Extra при росте — O(n). Известный handle исключает поиск; у list-insert-after без него надо сначала найти место. Ключи hash ограниченной длины, распределение хорошее; hash-set-growth обозначает один вызов с ростом массива buckets, n≈m.

Операции: list-find, list-insert-after, list-reverse, slice-index, slice-prepend, slice-stack-push, linked-queue-dequeue, ring-enqueue, hash-get, hash-set-growth. Для неизвестного имени вернуть поля unknown. Slice-prepend рассматривается с возможным ростом capacity. Пример: slice-index -> Time O(1), Extra O(1), Total O(cap).

Подсказка: выпишите предусловие каждой строки и отдельно ответьте, появляются ли новый узел, массив или только локальные указатели. 12 именованных тестов проверяют метки. Затем в `answers.md` объясните своими словами, почему амортизация не является worst-case гарантией.

## 07.2 SelectStructure

`func SelectStructure(scenario string) string`. Выберите метку по обязательным условиям:

| Сценарий | Обязательное условие |
|---|---|
| random-index / dense-indexes | частый произвольный доступ по плотному индексу |
| lifo-undo / pending-lifo | последнее добавленное действие должно выйти первым |
| fifo-unbounded | FIFO, динамический рост, O(1) операций при сохранённых head/tail |
| fifo-fixed-reject | фиксированная ёмкость, O(1) FIFO, full отклоняет новое |
| append-bytes | собрать весь byte-payload, растущий объём |
| lookup-key | прикладной lookup по строковому ключу стандартным Go-контейнером |
| remove-known-node-both-directions | удалить уже известный узел и идти по связям в обе стороны |
| reverse-chain | задача уже дана как конечная singly цепочка |
| retain-last-k | bounded поток, старые значения вытесняются |
| unknown | неизвестный контракт |

Допустимые метки перечислены в тестах: slice, stack, linked-queue, ring, buffer, map, doubly-list, singly-list, unknown. Подсказка: сначала выберите семантику, затем реализацию. Пример: неизвестное имя -> unknown. 12 сценариев. Для выбранной метки допишите в answers.md цену поиска и памяти, а также одно условие, при котором выбор поменялся бы.

Оба задания используют O(1) или O(длины имени) времени для фиксированной таблицы вариантов, O(1) дополнительной памяти. Тесты не проверяют литературный стиль объяснений: их преподаватель смотрит в PR.
