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
	@echo "B bin/gorecsel    ./cmd/gorecsel"
	@go build -o bin/gorecsel    ./cmd/gorecsel
install-go: all-go
	@install -d $(DESTDIR)$(PREFIX)/bin
	@echo I bin/gorecsel
	@cp bin/gorecsel    $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f bin/gorecsel
## -- AUTO-GO --
## -- license --
install: install-license
install-license: COPYING
	@echo 'I share/doc/go-recfile/COPYING'
	@mkdir -p $(DESTDIR)$(PREFIX)/share/doc/go-recfile
	@cp COPYING $(DESTDIR)$(PREFIX)/share/doc/go-recfile
## -- license --
