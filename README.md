# Домашняя работа 8 — структуры данных в Go

Продолжение алгоритмического блока после занятия 7. Go 1.22, стандартная библиотека, независимый модуль. **19 задач в 8 разделах** (design-класс с несколькими методами считается одной задачей). Все реализации оставлены TODO; helpers в tests строят входы и oracle, не содержат готовых решений.

## Разделы

| Документ | Задач | Реализация |
|---|---:|---|
| [docs/01_linkedlist.md](docs/01_linkedlist.md) | 5 | internal/linkedlist |
| [docs/02_stack.md](docs/02_stack.md) | 2 | internal/stack |
| [docs/03_queue.md](docs/03_queue.md) | 2 | internal/queue |
| [docs/04_buffer.md](docs/04_buffer.md) | 2 | internal/buffer |
| [docs/05_ringbuffer.md](docs/05_ringbuffer.md) | 2 | internal/ringbuffer |
| [docs/06_hashtable.md](docs/06_hashtable.md) | 3 | internal/hashtable |
| [docs/07_analysis.md](docs/07_analysis.md) | 2 | internal/analysis |
| [docs/08_benchmark.md](docs/08_benchmark.md) | 1 | internal/benchmark |

Восемь оригинальных LeetCode: 206, 21, 141, 20, 155, 232, 622, 217. Ссылки, краткий пересказ, сигнатуры, ограничения, подсказки и отличия API — в документах разделов. Методы empty/zero возвращают bool, чтобы хранить 0 и отрицательные без sentinel-неоднозначности. Контракты tests — часть задания.

## Работа и проверка

```bash
make help
make starter-check            # зелёный starter с TODO
make run-all                  # восемь примеров actual/expected
make test-linkedlist
make test-stack
make test-queue
make test-ringbuffer
make test-hashtable
make test-buffer
make test-analysis
make test-benchmark            # реальный вызов BenchmarkStructures: 1x
make test
make test-race
make coverage
make bench
make ci
```

Функциональные `test`, `test-race`, coverage и ci намеренно красные, пока TODO не реализованы. Benchmark тоже падает явно на 08.1. Не заменяйте проверки skip и не подставляйте ответы classroom. Нулевые/пустые случаи могут проходить даже у заглушки; остальные именованные сценарии выявляют отсутствие решения. Unit-тесты должны стать зелёными до измерений.

## Маршрут практики

Сначала Reverse/HasCycle/ValidParentheses и очередь через два стека. Затем MinStack/CircularQueue/ContainsDuplicate/Frequencies. Merge, собственная HashTable, ByteBuffer и остальные задачи — развёрнутая практика; все 19 не предполагаются за один вечер. 07_analysis и 08_benchmark завершают работу: объяснение, проверяемый контракт, реальные данные. Отправляйте PR с зелёными tests реализованных разделов, fmt/vet и своим коротким объяснением инвариантов/сложности. Общий CI станет зелёным после всех задач.

## Исполняемые примеры

| Команда | Свои функции, которые вызываются |
|---|---|
| `make run-linkedlist` | Reverse, HasCycle |
| `make run-stack` | ValidParentheses, MinStack |
| `make run-queue` | TwoStackQueue, RoundRobin |
| `make run-buffer` | ByteBuffer.Write/Bytes/Read |
| `make run-ringbuffer` | CircularQueue, RecentValues |
| `make run-hashtable` | HashTable, Frequencies, ContainsDuplicate |
| `make run-analysis` | Analyze, SelectStructure |
| `make run-benchmark` | Обнаружение своей BenchmarkStructures через `go test -list` |

Первые семь команд показывают **фактический** результат рядом с ожидаемым. У TODO выводится `TODO/ошибка`, а не подставленный правильный ответ. Сериализация результата Reverse ограничена восемью узлами; результат nil у фабрик проверяется до обращения к методам. Прохождение запуска не доказывает решение: контракт проверяют функциональные tests. Восьмая команда не измеряет время и не запускает тяжёлый benchmark при `run-all`; настоящий запуск — `make bench` после реализации 08.1.

Важное отличие Go: обычный `go test` не выполняет `Benchmark…`. Поэтому `make test-benchmark` / `make bench-check` выполняют отдельный smoke с `-benchtime=1x` и выявляют TODO08.1. `make ci` тоже включает эту проверку. До реализации задач starter зелёный, functional acceptance красный — это ожидаемое состояние, а не повод отключать tests.

## GitHub Actions и отправка решения

[Workflow и диагностика](docs/ci.md): push, PR и ручной запуск. Starter job проверяет компиляцию, формат, vet, восемь примеров и сборку. Отдельная матрица проверяет восемь разделов настоящими функциональными tests; работающие разделы видны отдельно. Не заменяйте их compilation-only проверкой.

Создайте ветку, реализуйте один раздел, запустите `make test-<section>` и приложите PR с объяснением инвариантов. Исправляйте первый упавший subtest из лога Actions. Все разделы, включая 08.1, должны стать зелёными перед полной сдачей. Покрытие — дополнительный сигнал, не замена корректности.

## Короткие видео и визуализации структур данных

Для самостоятельной практики без регистрации: [стек на массиве](https://www.cs.usfca.edu/~galles/visualization/StackArray.html), [очередь на списке](https://www.cs.usfca.edu/~galles/visualization/QueueLL.html), [кольцевой буфер](https://commons.wikimedia.org/wiki/File:Circular_Buffer_Animation.gif). Перед операцией предскажите следующий элемент и изменение ссылок/индексов. Затем запустите свой `cmd` и tests соответствующего раздела. Домашняя работа и CI не зависят от Instagram или доступа к видео.

## Если git push завершился HTTP 400

[Диагностика и проверка отправки](docs/git-push.md). `make git-check-push` сравнивает локальный HEAD с реальной веткой origin и ничего не отправляет. При необходимости повторите отправку разово через `git -c http.version=HTTP/1.1 push origin main`, затем проверьте SHA.
