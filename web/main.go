//go:build js && wasm
// +build js,wasm

package main

import (
	"aba/algorithms"
	"aba/components"
	jslayer "aba/js-layer"
	"context"
	"fmt"
	"strings"

	"syscall/js"
)

const TABLE_CONTAINER_ID = "table-container"

type TwoSumScreenHook struct{}

func (TwoSumScreenHook) Found([]int) {}

func (TwoSumScreenHook) OnStep(hashMap algorithms.HashMap, index int) {
	buf := new(strings.Builder)
	hashMapTable := components.HashMapTable(hashMap)
	hashMapTable.Render(context.Background(), buf)

	tableContainerDomElement := js.Global().Get("document").Call("getElementById", TABLE_CONTAINER_ID)
	tableContainerDomElement.Set("innerHTML", buf.String())
}

func twoSum(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		fmt.Println("You need two arguments min")
		return nil
	}

	if !args[0].InstanceOf(js.Global().Get("Array")) {
		fmt.Println("The first parameter must me an array of numbers")
		return nil
	}

	if args[1].Type() != js.TypeNumber {
		fmt.Println("The second parameter must me a number")
		return nil
	}

	input := jslayer.JsIntArray(args[0])
	target := args[1].Int()
	algorithms.TwoSum(input, target, TwoSumScreenHook{})

	// return jslayer.IntArrayToJs(result)

	// return js.ValueOf(buf.String())
	return nil
}

var counter = 0

func add(this js.Value, args []js.Value) any {

	counter++
	js.Global().
		Get("document").Call("getElementById", "output").
		Set("innerHTML", counter)

	return nil
}

func binaryTree(this js.Value, args []js.Value) any {
	input := jslayer.JsIntArray(args[0])
	tree := algorithms.CreateBinaryTree(input)
	componentHTML := new(strings.Builder)
	components.BinaryTree(tree).Render(context.Background(), componentHTML)

	fmt.Println(componentHTML)

	return js.ValueOf(componentHTML)
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("twoSum", js.FuncOf(twoSum))
	js.Global().Set("loop", js.FuncOf(algorithms.Loop))
	js.Global().Set("binaryTree", js.FuncOf(binaryTree))
}

func main() {
	registerCallbacks()
	fmt.Println("Golang server started.")

	js.Global().Set("TABLE_CONTAINER_ID", js.ValueOf(TABLE_CONTAINER_ID))

	components.AllAlgorithmsController()
	components.BynaryTreeController()

	select {}
}
