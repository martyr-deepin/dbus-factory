PREFIX  = /usr

GOPATH     := $(shell go env GOPATH)
QT5_LIBDIR ?= $(PREFIX)/lib/qt5
QMLDIR     ?= $(QT5_LIBDIR)/qml

JSONS  := $(wildcard *.in.json)
DBUS_GENERATOR = dbus-generator

define qml_pkg
$(shell basename $(1) .in.json | \
    sed -e 's:\(\.\)*\([a-zA-Z]*\):\1\u\2:g' \
        -e 's:\.:/:g')
endef

define qml_rule
QT_SRC_$(1) := $(addprefix qml/DBus/,$(call qml_pkg,$(1)))

$$(QT_SRC_$(1)):$(1)
	@echo [QML] $$(notdir $$@)
	@$(DBUS_GENERATOR) -in $$^ -out $$@ -target qml
	@make -C $$@
endef

define go_pkg
$(shell basename $(1) .in.json | tr A-Z a-z | tr '.' '/')
endef

define golang_rule
GO_PKG_$(1) := $(call go_pkg,$(1))
GO_SRC_$(1) := $(addprefix go/src/dbus/,$$(GO_PKG_$(1)))

$$(GO_SRC_$(1)):$(1)
	@echo [GOLANG] $$(notdir $$@)
	@$(DBUS_GENERATOR) -in $$^ -out $$@ -target golang
endef

GO_PKGS  := $(foreach i,$(JSONS),$(call go_pkg,$(i)))
GO_SRCS  := $(addprefix go/src/dbus/,$(GO_PKGS))
QT_PKGS  := $(foreach i,$(JSONS),$(call qml_pkg,$(i)))
QT_SRCS  := $(addprefix qml/DBus/,$(QT_PKGS))

all:          build
distclean:    clean

build:        build-golang   build-qml
install:      install-golang install-qml
build-qml:    $(QT_SRCS)
build-golang: $(GO_SRCS)

check:
	@true

clean:
	@rm -rf go qml

install-golang: build-golang
	@echo copying go source files to $(DESTDIR)$(GOPATH)/src ...
	@mkdir -p $(DESTDIR)$(GOPATH)/src
	@cp -r go/src/dbus $(DESTDIR)$(GOPATH)/src/

install-qml: build-qml
	@echo install qml plugins to $(DESTDIR)$(QMLDIR)
	@mkdir -p $(DESTDIR)$(QMLDIR)
	@for i in $(QT_SRCS); do \
          mkdir -p $(DESTDIR)$(QT5_LIBDIR)/$$i; \
          cp -r $$i/lib/* $(DESTDIR)$(QT5_LIBDIR)/$$i/; \
        done

$(foreach i,$(JSONS),$(eval $(call golang_rule,$(i))))
$(foreach i,$(JSONS),$(eval $(call qml_rule,$(i))))
