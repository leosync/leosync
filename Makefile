.PHONY: getdeps build build-multiarch doc check verifiers deadcode spelling fmt interfacer errcheck gocyclo lint run test vet

default: build

getdeps:
	@go get github.com/golang/lint/golint && echo "Installed golint:"
	@go get github.com/fzipp/gocyclo && echo "Installed gocyclo:"
	@go get github.com/remyoudompheng/go-misc/deadcode && echo "Installed deadcode:"
	@go get github.com/client9/misspell/cmd/misspell && echo "Installed misspell:"
	@go get github.com/mvdan/interfacer/cmd/interfacer && echo "Installed interfacer:"
	@go get github.com/kisielk/errcheck && echo "Installed errcheck:"

build: getdeps verifiers test
	go build -v -o ./bin/leosync .

build-multiarch: check
	go get github.com/karalabe/xgo
	docker pull karalabe/xgo-latest
	mkdir -p dist
	xgo -dest dist -go latest -v --targets='windows-6.1/amd64,windows-6.1/386,windows-6.1/arm-7,darwin-10.9/amd64,darwin-10.9/386,darwin-10.9/arm-7,linux/amd64,linux/386,linux/arm-7' github.com/igrybkov/leosync

doc:
	godoc -http=:6060 -index

check: getdeps verifiers test

verifiers: vet fmt lint gocyclo deadcode spelling errcheck interfacer

deadcode:
	@deadcode

spelling:
	@@find . -type f -name '*.go' -not -path "./vendor/*" | xargs -L1 misspell -error

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@find . -type f -name '*.go' -not -path "./vendor/*" | xargs -L1 gofmt -d -s

interfacer:
	@go list ./... | grep -vE '^vendor/' | interfacer

errcheck:
	@go list ./... | grep -v 'vendor/' | xargs -L1 errcheck -blank

gocyclo:
	@find . -iname '*.go' -not -path "./vendor/*" | xargs -L1 gocyclo -over 10

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	@find . -type f -name '*.go' -not -path "./vendor/*" | xargs -L1 golint -set_exit_status

run: build
	./bin/leosync

test:
	go test ./...

vendor_update:
	go get -u ./... && godep update ./...

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	@find . -type f -name '*.go' -not -path "./vendor/*" | xargs -L1 go tool vet
