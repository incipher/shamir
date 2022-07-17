install:
  go get

build:
  go build main.go

test:
  go clean -testcache
  CGO_ENABLED="0" go test -v ./...

publish version:
  @echo 'Publishing {{version}} ...'
  git tag -a {{version}} -m "{{version}}" -s
  git push origin {{version}}
  goreleaser release --rm-dist
