DESTDIR=
PREFIX=/usr/local
all:
clean:
install:
## -- AUTO-GO --
all:     all-go
install: install-go
clean:   clean-go
all-go:
	@echo "B bin/gorecsel$(EXE) ./cmd/gorecsel"
	@go build -o bin/gorecsel$(EXE) ./cmd/gorecsel
install-go: all-go
	@install -d $(DESTDIR)$(PREFIX)/bin
	@echo I bin/gorecsel$(EXE)
	@cp bin/gorecsel$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f bin/gorecsel
## -- AUTO-GO --
## -- license --
install: install-license
install-license: COPYING
	@echo 'I share/doc/go-recutils/COPYING'
	@mkdir -p $(DESTDIR)$(PREFIX)/share/doc/go-recutils
	@cp COPYING $(DESTDIR)$(PREFIX)/share/doc/go-recutils
## -- license --
