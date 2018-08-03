# print test

> test
```
cd $GOPATH/src/go-test-demo/print && \
go test -v -cover=true -coverprofile=/tmp/print-test-cover.out ./... && \
go tool cover -html=/tmp/print-test-cover.out -o=/tmp/print-test.html && \
open /tmp/print-test.html
```