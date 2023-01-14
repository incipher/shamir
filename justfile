set dotenv-load

install-dependencies:
  go get

upgrade-dependencies:
  go get -u
  go mod tidy

upgrade-go-version version:
  go mod edit -go {{version}}
  go mod tidy

test:
  go clean -testcache
  CGO_ENABLED="0" go test -v ./...

build:
  go build main.go

publish version:
  @echo 'Publishing {{version}} ...'
  git tag -a {{version}} -m "{{version}}" -s
  git push origin {{version}}
  goreleaser release --rm-dist
