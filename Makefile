SHELL := /bin/bash
.DEFAULT_GOAL := help
GO ?= go
BENCHTIME ?= 100ms
BENCHCOUNT ?= 3
BENCHCPU ?= 1
COVERAGE_FILE ?= coverage.out

fmt: ## КАЧЕСТВО | Применить gofmt
	gofmt -w $$(find . -name '*.go' -not -path './bin/*')
fmt-check: ## КАЧЕСТВО | Проверить форматирование без изменения файлов
	@files="$$(gofmt -l $$(find . -name '*.go' -not -path './bin/*'))"; if [ -n "$$files" ]; then echo "$$files"; exit 1; fi
vet: ## КАЧЕСТВО | Статический анализ go vet
	$(GO) vet ./...
compile: ## КАЧЕСТВО | Компиляция всех пакетов и тестов без выполнения заданий
	$(GO) test -run '^$$' ./...
test-unit: ## ТЕСТЫ | Все модульные тесты internal
	$(GO) test ./internal/...
test-race: ## ТЕСТЫ | Все тесты с race detector
	$(GO) test -race ./...
coverage: ## КАЧЕСТВО | Statements coverage всех internal-пакетов
	$(GO) test ./internal/... -coverpkg=./internal/... -covermode=atomic -coverprofile=$(COVERAGE_FILE)
	$(GO) tool cover -func=$(COVERAGE_FILE)
clean: ## СБОРКА | Удалить bin и coverage.out
	rm -rf bin $(COVERAGE_FILE)

SECTIONS := linkedlist stack queue buffer ringbuffer hashtable analysis benchmark
CMDS := 01_linkedlist 02_stack 03_queue 04_buffer 05_ringbuffer 06_hashtable 07_analysis 08_benchmark
BENCHFLAGS := -benchmem -benchtime=$(BENCHTIME) -count=$(BENCHCOUNT) -cpu=$(BENCHCPU)

help: ## СПРАВКА | Все команды; 19 задач остаются TODO
	@awk 'BEGIN { FS=":.*## " } /^[a-zA-Z0-9_-]+:.*## / { split($$2, p, "[|]"); g=p[1]; sub(/ +$$/, "", g); sub(/^ +/, "", p[2]); if (!(g in known)) { order[++n]=g; known[g]=1 } text[g]=text[g] sprintf("  make %-25s %s\n", $$1, p[2]) } END { for (i=1;i<=n;i++) printf "\n%s\n%s", order[i], text[order[i]] }' $(MAKEFILE_LIST)
	@printf '\nrun показывает факт; test проверяет контракт. Starter зелёный; functional tests до решений красные.\n'
	@printf 'make run-linkedlist; make test-linkedlist; make test-benchmark — 1x проверка лаборатории.\n'
	@printf 'make bench BENCHTIME=100ms BENCHCOUNT=3 BENCHCPU=1 — реальные измерения после решений.\n'

starter-check: compile fmt-check vet test-demos build ## ПРОВЕРКИ STARTER | Компиляция, стиль, vet, восемь cmd и сборка

test: test-unit test-integration ## ТЕСТЫ | Все функциональные тесты и проверки cmd

test-integration: ## ТЕСТЫ | Проверка запуска восьми cmd
	$(GO) test ./test/integration/...
test-demos: ## ПРОВЕРКИ STARTER | Формат вывода и запуск cmd, без проверки TODO
	$(GO) test ./internal/demo ./test/integration/...

test-linkedlist: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты linkedlist; до решения ожидаются ошибки
test-stack: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты stack; до решения ожидаются ошибки
test-queue: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты queue; до решения ожидаются ошибки
test-buffer: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты buffer; до решения ожидаются ошибки
test-ringbuffer: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты ringbuffer; до решения ожидаются ошибки
test-hashtable: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты hashtable; до решения ожидаются ошибки
test-analysis: ## ТЕСТЫ ПО РАЗДЕЛАМ | Контракты analysis; до решения ожидаются ошибки
$(addprefix test-,$(filter-out benchmark,$(SECTIONS))):
	$(GO) test ./internal/$(patsubst test-%,%,$@)

test-benchmark: ## ТЕСТЫ ПО РАЗДЕЛАМ | Реально вызвать TODO 08.1 один раз; это не замер скорости
	$(GO) test ./internal/benchmark -run '^$$' -bench '^BenchmarkStructures$$' -benchmem -benchtime=1x -count=1 -cpu=1
bench-check: test-benchmark ## ТЕСТЫ | Псевдоним проверки 19-й задачи
bench: ## БЕНЧМАРКИ | Измерения самостоятельной лаборатории; TODO приводит к ошибке
	$(GO) test ./internal/benchmark -run '^$$' -bench '^BenchmarkStructures$$' $(BENCHFLAGS)

build: ## СБОРКА | Собрать восемь cmd в bin
	@mkdir -p bin
	@set -e; for cmd in $(CMDS); do $(GO) build -buildvcs=false -o bin/$$cmd ./cmd/$$cmd; done
ci: compile fmt-check vet test test-race coverage bench-check build ## КАЧЕСТВО | Полное решение всех 19 задач, включая benchmark 08.1

run-all: ## ЗАПУСК | Восемь примеров; benchmark только discovery, без измерений
	@set -e; for cmd in $(CMDS); do $(GO) run ./cmd/$$cmd; done
run-linkedlist: ## ЗАПУСК | Демонстрация linkedlist
	$(GO) run ./cmd/01_linkedlist
run-stack: ## ЗАПУСК | Демонстрация stack
	$(GO) run ./cmd/02_stack
run-queue: ## ЗАПУСК | Демонстрация queue
	$(GO) run ./cmd/03_queue
run-buffer: ## ЗАПУСК | Демонстрация buffer
	$(GO) run ./cmd/04_buffer
run-ringbuffer: ## ЗАПУСК | Демонстрация ringbuffer
	$(GO) run ./cmd/05_ringbuffer
run-hashtable: ## ЗАПУСК | Демонстрация hashtable
	$(GO) run ./cmd/06_hashtable
run-analysis: ## ЗАПУСК | Демонстрация analysis
	$(GO) run ./cmd/07_analysis
run-benchmark: ## ЗАПУСК | Демонстрация benchmark
	$(GO) run ./cmd/08_benchmark

git-check-push: ## GIT | Сравнить HEAD с веткой origin, без отправки и изменения refs
	@set -e; \
	branch="$$(git symbolic-ref --quiet --short HEAD)" || { echo "Нужна текущая ветка (HEAD detached)." >&2; exit 1; }; \
	local_sha="$$(git rev-parse HEAD)"; \
	remote_line="$$(git ls-remote --exit-code origin "refs/heads/$$branch")" || { echo "Не удалось проверить ветку на origin." >&2; exit 1; }; \
	remote_sha="$$(printf '%s\n' "$$remote_line" | awk '{print $$1}')"; \
	printf 'Ветка: %s\nHEAD:   %s\norigin: %s\n' "$$branch" "$$local_sha" "$$remote_sha"; \
	if [ "$$local_sha" = "$$remote_sha" ]; then echo "SHA совпадают: эта вершина есть на origin."; \
	else echo "SHA различаются. Проверьте историю и результат push." >&2; exit 1; fi

.PHONY: fmt fmt-check vet compile test-unit test-race coverage clean help starter-check test test-integration test-demos test-linkedlist test-stack test-queue test-buffer test-ringbuffer test-hashtable test-analysis test-benchmark bench-check bench build ci run-all run-linkedlist run-stack run-queue run-buffer run-ringbuffer run-hashtable run-analysis run-benchmark git-check-push
