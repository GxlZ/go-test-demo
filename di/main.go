package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/facebookgo/inject"
)

type HomePlanetRenderApp struct {
	NameAPI   *NameAPI   `inject:""`
	PlanetAPI *PlanetAPI `inject:""`
}

func (a *HomePlanetRenderApp) Render(id uint64) string {
	return fmt.Sprintf(
		"%s is from the planet %s.",
		a.NameAPI.Name(id),
		a.PlanetAPI.Planet(id),
	)
}

type NameAPI struct {
	HTTPTransport http.RoundTripper `inject:""`
}

func (n *NameAPI) Name(id uint64) string {
	return "Spock"
}

type PlanetAPI struct {
	HTTPTransport http.RoundTripper `inject:""`
}

func (p *PlanetAPI) Planet(id uint64) string {
	return "Vulcan"
}

func main() {
	var g inject.Graph
	var a HomePlanetRenderApp
	err := g.Provide(
		&inject.Object{Value: &a},
		&inject.Object{Value: http.DefaultTransport},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(a.Render(42))

}
