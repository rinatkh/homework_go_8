# 04. buffer — 2 задач

Функциональные тесты находятся рядом с TODO в internal/. Решения здесь не приведены. nil и пустой результат эквивалентны по элементам, если явно не требуется handle.

## 04.1 ByteBuffer

```go
type ByteBuffer; Write([]byte)(int,error); Read([]byte)(int,error); Bytes()[]byte; Len()int; Reset()
```

**Контракт:** Линейный byte-buffer. Write копирует, Read потребляет, Bytes — независимая копия, Reset обнуляет used область и сохраняет cap.

**Пример:** `Write(Go API),Read(3): Go ; unread API`.

**Ограничения:** Общий ввод≤1 MiB; без bytes.Buffer как реализации. Read(nil) -> 0,nil; empty non-empty-read -> 0,io.EOF. Это контракт байтов, не rune.

**Подсказка:** Отделите физическую длину data от unread; сравните alias-контракт std Bytes.

**Ожидаемая сложность:** Read k O(k), Write k amortized O(k) с ростом/переносом worst O(n+k), Bytes O(n); total O(cap).

**Тесты:** 13 размеров, alias, EOF, нулевой Read, Reset и сохранение cap.

## 04.2 AppendBounded

```go
func AppendBounded(dst, chunk []byte, limit int) ([]byte,bool)
```

**Контракт:** Атомарно добавить chunk, если итоговая длина≤limit. При отказе вернуть прежний dst без записи.

**Пример:** `ab + cd,limit=3 -> ab,false; limit=4 -> abcd,true`.

**Ограничения:** limit<0 -> отказ, len(dst)>limit тоже отказ; alias dst разрешён, chunk не меняется.

**Подсказка:** Проверка должна выполняться до первой записи, включая пустой chunk.

**Ожидаемая сложность:** O(k) без роста, worst O(n+k), extra O(n+k) при росте.

**Тесты:** 13 сценариев, exact/full/negative, Unicode как байты, unchanged при отказе.
