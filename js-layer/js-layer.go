//go:build js && wasm
// +build js,wasm

package jslayer

import (
	"syscall/js"
)

// IntArray gets
func JsIntArray(arr js.Value) []int {
	result := make([]int, arr.Length())
	for i := 0; i < arr.Length(); i++ {
		result[i] = arr.Index(i).Int()
	}
	return result
}

func IntArrayToJs(arr []int) js.Value {
	jsResult := js.Global().Get("Array").New()
	for _, num := range arr {
		jsResult.Call("push", num)
	}

	return jsResult
}
