//go:build js && wasm
// +build js,wasm

package algorithms

import (
	"syscall/js"
	"time"
)

func Loop(this js.Value, args []js.Value) any {
	go func() {
		for i := 1; i <= 10; i++ {
			// Get the HTML div element by its ID
			div := js.Global().Get("document").Call("getElementById", "output")

			// Set the inner HTML to the current index
			div.Set("innerHTML", i)

			// Wait for 1 second before the next update
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}
