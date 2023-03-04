MODES := dief dieNil die dieExit diefExit dieExitNil
TEST  := testdata
DIE   := $(TEST)/die

.PHONY: $(DIE)

all: setup $(DIE)
	diff -u $(DIE) expect.die.log

setup:
	mkdir -p $(TEST)
	rm -f $(TEST)/*

$(DIE): $(MODES)

$(MODES):
	@go run die.go -mode $@ 2>> $(DIE) || (echo "$@:$$?" 2>> $(DIE))
