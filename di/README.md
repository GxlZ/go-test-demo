# redis test

> test
```
cd $GOPATH/src/go-test-demo/di && \
go test -v -cover=true -coverprofile=/tmp/di-test-cover.out ./... && \
go tool cover -html=/tmp/di-test-cover.out -o=/tmp/di-test.html && \
open /tmp/di-test.html
```