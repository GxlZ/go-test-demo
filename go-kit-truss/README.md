# go-kit-truss test

> test
```
cd $GOPATH/src/go-test-demo/go-kit-truss && \
go test -v -cover=true -coverprofile=/tmp/go-kit-truss-test-cover.out ./... && \
go tool cover -html=/tmp/go-kit-truss-test-cover.out -o=/tmp/go-kit-truss-test.html && \
open /tmp/go-kit-truss-test.html
```