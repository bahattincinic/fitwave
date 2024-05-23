default: vet build

all: vet test fitwave

build: fitwave

fmt:
	find . ! -path "*/vendor/*" -type f -name '*.go' -exec gofmt -l -s -w {} \;

test:
	go test -race -v ./...

swaggen:
	swag fmt && swag init --parseDependency -g ./cmd/fitwave/main.go -o ./api/docs

testci: test

vet:
	go vet ${GCFLAGS} ./...

fitwave:
	go build ${GCFLAGS} -ldflags "${LDFLAGS}" ./cmd/fitwave

clean:
	rm -vf ./fitwave

.PHONY: all fmt build clean vet fitwave

.NOTPARALLEL:
