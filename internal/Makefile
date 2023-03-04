TEST  := $(PWD)/testdata
TESTS := die stdin

.PHONY: $(TESTS)

all: setup $(TESTS)

setup:
	mkdir -p $(TEST)
	rm -f $(TEST)/*

$(TESTS):
	make -C $@ TESTDATA=$(TEST)/$@
	@diff -u $(TEST)/$@ $@/expect.log
