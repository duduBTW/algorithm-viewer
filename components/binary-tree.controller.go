// //go:build js && wasm
// // +build js,wasm

// package components

// import (
// 	"aba/algorithms"
// 	"context"
// 	"fmt"
// 	"math"
// 	"strconv"
// 	"strings"
// 	"syscall/js"
// )

// var nodes = []int{}

// func BynaryTreeController() {
// 	// Draw lines when app first opens
// 	DrawLines()
// 	js.Global().Get("window").Call("addEventListener", "resize", resizeHandler)

// 	// Add node logic
// 	form := js.Global().Get("document").Call("getElementById", CreateBinaryTreeIds.Form)
// 	if form.Type() != js.TypeNull {
// 		form.Call("addEventListener", "submit", addNodeFormSubmitHandler)
// 	}

// 	// Clear button
// 	clearButton := js.Global().Get("document").Call("getElementById", CreateBinaryTreeIds.ClearButton)
// 	if clearButton.Type() != js.TypeNull {
// 		clearButton.Call("addEventListener", "click", clearClickHandler)
// 	}
// }

// var clearClickHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 	// Update nodes
// 	nodes = []int{}

// 	// Renders tree
// 	RenderTree()
// 	return nil
// })

// var resizeHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 	// Draw lines on windown resize
// 	DrawLines()
// 	return nil
// })

// var addNodeFormSubmitHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 	event := args[0]
// 	// Stop form from reloading the page
// 	event.Call("preventDefault")

// 	input := js.Global().Get("document").Call("getElementById", CreateBinaryTreeIds.AddInput)
// 	inputValue := input.Get("value").String()

// 	value, err := strconv.Atoi(inputValue)
// 	if err != nil {
// 		fmt.Println("Typed value is not a valid number.")
// 		return nil
// 	}

// 	// Clears value
// 	input.Set("value", "")

// 	// Update nodes
// 	nodes = append(nodes, value)

// 	// Renders tree
// 	RenderTree()

// 	return nil
// })

// func RenderTree() {
// 	// Sets html
// 	treeContainer := js.Global().Get("document").Call("getElementById", CreateBinaryTreeIds.TreeContainer)
// 	binaryTreeHTML := new(strings.Builder)
// 	BinaryTree(algorithms.CreateBinaryTree(nodes)).Render(context.Background(), binaryTreeHTML)
// 	treeContainer.Set("innerHTML", binaryTreeHTML.String())

// 	// Draw lines
// 	DrawLines()
// }

// func DrawLines() {
// 	lines := js.Global().Get("document").Call("querySelectorAll", "#"+BinaryTreeIds.NodeLine)
// 	for i := 0; i < lines.Length(); i++ {
// 		line := lines.Index(i)
// 		DrawLine(line)
// 	}
// }

// func DrawLine(line js.Value) {
// 	start := line.Get("dataset").Get("start").String()
// 	end := line.Get("dataset").Get("end").String()

// 	startNode := NodeQuerySelector(start)
// 	endNode := NodeQuerySelector(end)

// 	startRect := startNode.Call("getBoundingClientRect")
// 	endRect := endNode.Call("getBoundingClientRect")

// 	startX := startRect.Get("left").Float() + startRect.Get("width").Float()/2
// 	startY := startRect.Get("bottom").Float()
// 	endX := endRect.Get("left").Float() + endRect.Get("width").Float()/2
// 	endY := endRect.Get("top").Float()

// 	// Calculate the angle and length
// 	deltaX := endX - startX
// 	deltaY := endY - startY
// 	angle := math.Atan2(deltaY, deltaX) * (180 / math.Pi) // Angle in degrees
// 	length := math.Sqrt(deltaX*deltaX + deltaY*deltaY)

// 	// Set CSS properties for positioning and rotating the line
// 	line.Get("style").Call("setProperty", "--start-x", fmt.Sprintf("%fpx", startX))
// 	line.Get("style").Call("setProperty", "--start-y", fmt.Sprintf("%fpx", startY))
// 	line.Get("style").Call("setProperty", "--line-angle", fmt.Sprintf("%fdeg", angle))
// 	line.Get("style").Call("setProperty", "--line-length", fmt.Sprintf("%fpx", length))

// 	fmt.Println("Connecting:", startNode, "to", endNode, "with angle:", angle, "and length:", length)
// }

// func NodeQuerySelector(value string) js.Value {
// 	selector := fmt.Sprintf("#%s[data-val=\"%s\"]", BinaryTreeIds.NodeVal, value)
// 	return js.Global().Get("document").Call("querySelector", selector)
// }
