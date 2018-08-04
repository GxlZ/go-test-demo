package di

import (
	"go.uber.org/dig"
	"go-test-demo/go-kit-truss/handlers"
)

var Container *dig.Container

func init() {
	Container = dig.New()
	Container.Provide(NewRedisConn)
	Container.Provide(handlers.NewService)
}