APP?=advent
PORT?=8000

clean:
	rm -f ${APP}

build: clean
	go build -o ${APP}

run: build
	PORT=${PORT} ./${APP}

test:
	go test -v -race ./...
