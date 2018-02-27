

all: test


test: go test -v ./... -cover -bench . -race