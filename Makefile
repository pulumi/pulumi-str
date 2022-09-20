
GEN     := ./bin/pulumi-gen-str
VERSION := 0.1.0

build:
	mkdir -p bin
	cd str && go build \
		-o ../bin \
		-ldflags "-X github.com/pulumi/pulumi-str/str/version.Version=${VERSION}" ./...

tidy:
	cd str && go mod tidy

sdk_prep: build
	mkdir -p sdk
build_sdks: build build_dotnet_sdk build_nodejs_sdk build_python_sdk build_go_sdk
	if [ -f sdk/go.mod ]; then rm sdk/go.mod; fi
	cd sdk && go mod init github.com/pulumi/pulumi-str/sdk
	${GEN} schema | jq > sdk/schema.json

build_%_sdk: sdk_prep
	if [ -d sdk/$* ]; then rm -rf sdk/$*; fi
	${GEN} language $*

export PULUMI_CONFIG_PASSPHRASE := "not-secret"
test_prep: build_sdks
	mkdir $$PWD/state
	pulumi login --cloud-url file://$$PWD/state
test: test_go test_nodejs
	pulumi logout
	rm -r $$PWD/state
	if [[ "$$CI" == "" ]]; then pulumi login; fi
# If we are running locally, we log back in to the previous pulumi account.
# We don't do this in CI since the CI has no previous account.

test_go: test_prep
	cd tests/golang && pulumi stack select -c test_go
	cd tests/golang && pulumi up --yes --skip-preview
	cd tests/golang && pulumi destroy --yes --skip-preview
	cd tests/golang && pulumi stack rm --yes

test_nodejs: test_prep
	cd sdk/nodejs && yarn link && yarn install
	cd tests/nodejs && yarn link "@pulumi/str"
	cd tests/nodejs && yarn install
	cd tests/nodejs && pulumi stack select -c test_nodejs
	cd tests/nodejs && pulumi up --yes --skip-preview
	cd tests/nodejs && pulumi destroy --yes --skip-preview
	cd tests/nodejs && pulumi stack rm --yes

lint: lint-golang lint-copyright
lint-golang:
	cd str && golangci-lint run -c ../.golangci.yml --timeout 5m
lint-copyright:
	pulumictl copyright -x 'examples/**' -x 'sdk/**'
