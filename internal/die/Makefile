MODES   := dief dieNil die dieExit diefExit dieExitNil

all: $(MODES)

$(MODES):
	@go run main.go -mode $@ 2>> $(TESTDATA) || (echo "$@:$$?" 2>> $(TESTDATA))
