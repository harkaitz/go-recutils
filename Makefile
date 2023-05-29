DESTDIR=
PREFIX=/usr/local
all:
clean:
install:
## -- AUTO-GO --
GO_PROGRAMS += bin/gorecsel$(EXE) 
.PHONY all-go: $(GO_PROGRAMS)
all:     all-go
install: install-go
clean:   clean-go
deps:
bin/gorecsel$(EXE): deps 
	go build -o $@ $(GORECSEL_FLAGS) $(GO_CONF) ./cmd/gorecsel
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp bin/gorecsel$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(GO_PROGRAMS)
## -- AUTO-GO --
## -- license --
install: install-license
install-license: COPYING
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/go-recutils
	cp COPYING $(DESTDIR)$(PREFIX)/share/doc/go-recutils
## -- license --
## -- AUTO-SERVICE --

## -- AUTO-SERVICE --
