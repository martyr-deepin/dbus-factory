all: check build

in.json/%-test:
	@jq . in.json/$* 1>/dev/null || (echo "==> invalid json:" "$*"; exit 100)

xml/%-test:
	@xmllint --noout xml/$* || (echo "==> invalid xml:" "$*"; exit 100)

check: $(foreach json, $(wildcard in.json/*.json), $(json)-test) $(foreach xml, $(wildcard xml/*.xml), $(xml)-test)


%:
	@+make -C in.json $@

.PHONY: check
