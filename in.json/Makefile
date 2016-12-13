PREFIX  = /usr

GOPATH     := $(shell go env GOPATH)

JSONS  := $(wildcard *.in.json)
DBUS_GENERATOR = dbus-generator

define go_pkg
$(shell basename $(1) .in.json | tr A-Z a-z | tr '.' '/')
endef

define golang_rule
$(shell $(DBUS_GENERATOR) -in $(1) -out $(addprefix go/src/dbus/,$(call go_pkg,$(1))) -target golang)
endef

GO_PKGS  := $(foreach i,$(JSONS),$(call go_pkg,$(i)))
GO_SRCS  := $(addprefix go/src/dbus/,$(GO_PKGS))

all:          build
distclean:    clean

build:        build-golang
install:      install-golang
build-golang: $(GO_SRCS)

check:
	@true

clean:
	@rm -rf go

install-golang: build-golang
	@echo copying go source files to $(DESTDIR)$(GOPATH)/src ...
	@mkdir -p $(DESTDIR)$(GOPATH)/src
	@cp -r go/src/dbus $(DESTDIR)$(GOPATH)/src/


$(foreach i,$(JSONS),$(eval $(call golang_rule,$(i))))
