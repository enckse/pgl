all: unittests checks

unittests:
	go test ./...

checks:
	make -C internal/
