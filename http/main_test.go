package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"bytes"
	"net/http/httptest"
	"net/http"
)

func TestGreet(t *testing.T) {

	Convey("hello method", t, func() {
		Convey("input name: go-kit", func() {
			name := "go-kit"
			b := bytes.Buffer{}
			hello(&b, name)
			So(b.String(), ShouldEqual, "hello,go-kit")
		})

		Convey("input name: empty", func() {
			b := bytes.Buffer{}
			hello(&b, "")
			So(b.String(), ShouldEqual, "hello,world")
		})
	})

}

func TestHelloHandler(t *testing.T) {
	Convey("hello http handler: 'get /'", t, func() {
		req := httptest.NewRequest("GET", "/", nil)
		recorder := httptest.NewRecorder()
		helloHandler(recorder, req)

		Convey("status code", func() {
			So(recorder.Code, ShouldEqual, http.StatusOK)
		})

		Convey("body", func() {
			So(recorder.Body.String(), ShouldEqual, "hello,world")
		})
	})

}
