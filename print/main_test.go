package main

import (
	"testing"
	//"fmt"
	//"github.com/stretchr/testify/mock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {

	Convey("hello method", t, func() {
		Convey("input name: go-kit", func() {
			name := "go-kit"
			got := hello(name)
			want := helloPrefix + name
			So(got, ShouldEqual, want)
		})

		Convey("input name: empty", func() {
			name := ""
			got := hello(name)
			want := helloPrefix + "world"
			So(got, ShouldEqual, want)
		})
	})

}
