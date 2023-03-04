MODES := one multiple

all: $(MODES)

$(MODES):
	printf "one\ntwo" | go run main.go -mode $@ >> $(TESTDATA)
