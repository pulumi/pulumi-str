
GEN := ./bin/pulumi-gen-str

build:
	mkdir -p bin
	cd str && go build -o ../bin ./...

tidy:
	cd str && go mod tidy

build_sdks: build
	mkdir -p sdk
	${GEN} schema | jq > sdk/schema.json
	${GEN} language nodejs
	${GEN} language python
	${GEN} language go
	${GEN} language dotnet
