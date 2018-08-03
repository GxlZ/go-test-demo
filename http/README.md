# http handler test

> test
```
cd $GOPATH/src/go-test-demo/http && \
go test -v -cover=true -coverprofile=/tmp/http-test-cover.out ./... && \
go tool cover -html=/tmp/http-test-cover.out -o=/tmp/http-test.html && \
open /tmp/http-test.html
```