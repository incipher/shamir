install:
  go get

build:
  go build main.go

test:
  go clean -testcache
  CGO_ENABLED="0" go test -v ./...
