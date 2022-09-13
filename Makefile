
build:
	mkdir -p bin
	cd str && go build -o ../bin ./...

tidy:
	cd str && go mod tidy
