# CI — Homework 8

Workflow: [.github/workflows/ci.yml](../.github/workflows/ci.yml). Триггеры — push, pull_request, workflow_dispatch; ручной запуск: Actions → Homework 8 CI → Run workflow. `go.mod` задаёт Go 1.22; secrets и дополнительные сервисы не нужны.

Две независимые проверки:

| Проверка | Команды | Начальное состояние |
|---|---|---|
| Starter validation | `make starter-check`; `go test -race ./internal/demo ./test/integration/...` | Зелёное: TODO компилируются, восемь cmd запускаются |
| Functional tests, 8 разделов | `make test-linkedlist` … `make test-benchmark` | Красное до реализации соответствующих задач |

Все 19 исходных задач и tests сохранены. Группа benchmark содержит только benchmark-функцию, поэтому её acceptance-команда реально вызывает `BenchmarkStructures` один раз (`-benchtime=1x`); `go test -run '^$'` этого не делает. `08_benchmark` в cmd выполняет только поиск имени benchmark и не является проверкой решения 08.1.

При готовом разделе дополнительно запускаются race и coverage. Отчёт покрытия сохраняется в artifact на 7 дней, если он создан; провал acceptance не маскируется. Матрица `fail-fast: false` даёт увидеть результаты остальных разделов. В workflow нет `continue-on-error`, подавления кода возврата или пропуска tests ради зелёного статуса.

Локально:

```bash
make starter-check
make test-linkedlist
make test-benchmark
make ci
```

Первые команды нужны на каждом этапе. `make ci` ожидаемо красный на стартовых TODO; после всех реализаций он должен пройти и включает benchmark smoke. `make bench` запускайте для реальных измерений после correctness tests; его результат не сравнивается с порогом скорости в CI.

Для PR создайте ветку, запушьте её и откройте pull request. На вкладке Checks выберите красный раздел, затем step проверки контрактов: исправляйте первую ошибку со ссылкой на имя subtest. Starter зелёный и один раздел зелёный не означают полную сдачу. При ограничениях Actions в fork владелец может подтвердить стандартный запуск PR. Настройка branch protection остаётся в настройках репозитория.
