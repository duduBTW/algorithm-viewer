//go:build js && wasm
// +build js,wasm

package components

import (
	"context"
	"strings"
	"syscall/js"

	"github.com/a-h/templ"
)

type ComponentRegistry struct {
	cleanUpFns []*func()
}

var GlobalComponentRegistry = &ComponentRegistry{}

func (componentRegistry *ComponentRegistry) Register(cleanUpFn *func()) {
	componentRegistry.cleanUpFns = append(componentRegistry.cleanUpFns, cleanUpFn)
}

type BaseComponent struct {
	Target    string
	Component templ.Component
	SetupFn   func(reRender func()) func()
}

func (base BaseComponent) Render() {
	target := js.Global().Get("document").Call("querySelector", base.Target)

	var reRender = func() {
		componentHTML := new(strings.Builder)
		base.Component.Render(context.Background(), componentHTML)

		target.Set("innerHTML", componentHTML.String())
	}

	reRender()
	cleanUp := base.SetupFn(reRender)
	GlobalComponentRegistry.Register(&cleanUp)
}
