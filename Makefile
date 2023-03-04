INTERNALS := $(dir internal/**/Makefile)

.PHONY: $(INTERNALS)

all: unittests $(INTERNALS)

unittests:
	go test ./...

$(INTERNALS):
	make -C $@
