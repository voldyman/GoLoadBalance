
BIN = balance
SOURCES = src/main.go src/backends.go src/listener.go

all:
	go build  -o $(BIN) $(SOURCES)
