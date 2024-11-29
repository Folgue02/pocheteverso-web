init:
	if [ ! -d ./bin ]; then \
		mkdir bin; \
	fi

run:
	go run .

build:
	go build -o bin/pocheteverso . 

dev:
	go run . -port 8080 -static ./static

build-scripts: init
	go build -o bin/pv-install scripts/install.go