
GEN := ./bin/pulumi-gen-str

build:
	mkdir -p bin
	cd str && go build -o ../bin ./...

tidy:
	cd str && go mod tidy

build_sdks: build
	rm -r sdk
	mkdir -p sdk
	cd sdk && go mod init
	${GEN} schema | jq > sdk/schema.json
	${GEN} language nodejs
	${GEN} language python
	${GEN} language go
	${GEN} language dotnet

test_prep: build_sdks
test: test_go
test_go: test_prep
	cd tests/golang && go build ./...
	cd tests/golang && pulumi stack select -c dev
	cd tests/golang && pulumi up --yes --skip-preview
	cd tests/golang && pulumi destroy --yes --skip-preview
	cd tests/golang && pulumi stack rm --yes
