all: unittests checks vers

vers:
	go run internal/version/main.go version/vers.txt

unittests:
	go test ./...

checks:
	make -C internal/
