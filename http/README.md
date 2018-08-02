# go-tdd-demo

> test
```
cd $GOPATH/src/go-tdd-demo/mock
go test -v -cover=true -coverprofile=/tmp/test-cover.out ./... && \
go tool cover -html=/tmp/test-cover.out -o=/tmp/test-cover.html && \
open /tmp/test-cover.html
```