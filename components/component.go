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
	components []*BaseComponent
}

var GlobalComponentRegistry = &ComponentRegistry{}

func (componentRegistry *ComponentRegistry) Register(component *BaseComponent) {
	componentRegistry.components = append(componentRegistry.components, component)
}

type BaseComponent struct {
	Target    string
	Component templ.Component
	CleanupFn func()
	SetupFn   func()
}

func (base BaseComponent) Render() {
	target := js.Global().Get("document").Call("querySelector", base.Target)

	componentHTML := new(strings.Builder)
	base.Component.Render(context.Background(), componentHTML)

	target.Set("innerHTML", componentHTML.String())

	base.SetupFn()
	GlobalComponentRegistry.Register(&base)
}
