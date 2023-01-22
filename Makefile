.SILENT:

build:
	go mod verify
	go mod tidy
	go build -o .bin/calc cmd/main.go

run: build
	./.bin/calc

prod: build
	cp ./.bin/calc ${HOME}/.config/scripts/bin/calc
