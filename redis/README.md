# redis test

> test
```
cd $GOPATH/src/go-test-demo/redis && \
go test -v -cover=true -coverprofile=/tmp/redis-test-cover.out ./... && \
go tool cover -html=/tmp/redis-test-cover.out -o=/tmp/redis-test.html && \
open /tmp/redis-test.html
```