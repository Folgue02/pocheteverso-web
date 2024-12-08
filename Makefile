TARGET_BIN=bin/pvw

init:
	if [ ! -d ./bin ]; then \
		mkdir bin; \
	fi

run:
	go run .

build:
	go build -o bin/pvw . 

dev: build
	$(TARGET_BIN) -port 8080 -static ./static -assets ./assets -dynres ./dynres

build-scripts: init
	go build -o bin/pv-install scripts/install.go
