all: build

%:
	@+make -C in.json $@
