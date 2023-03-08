all: unittests checks

unittests:
	go test ./...

checks:
	cd internal && make
