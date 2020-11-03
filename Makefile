install:
	GO114MODULE=on go mod tidy
build: install
	GO114MODULE=on go build -o redis-grid.bin .

run: build
	./redis-grid.bin
