//go:build js && wasm
// +build js,wasm

package components

import (
	"context"
	"slices"
	"strings"
	"syscall/js"
)

func AllAlgorithmsController() {
	clickHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		rawPageId := event.Get("currentTarget").Get("dataset").Get("pageid")
		if rawPageId.Type() != js.TypeString {
			return nil
		}

		pageId := rawPageId.String()
		targetPageIndex := slices.IndexFunc(ClientAllAlgorithmsPages, func(page Page) bool { return page.Id == pageId })
		if targetPageIndex == -1 {
			return nil
		}

		for _, cleanUpFnPointer := range GlobalComponentRegistry.cleanUpFns {
			cleanUpFn := *cleanUpFnPointer
			cleanUpFn()
		}

		targetPage := ClientAllAlgorithmsPages[targetPageIndex]

		componentHTML := new(strings.Builder)
		targetPage.Component().Render(context.Background(), componentHTML)

		js.Global().Get("document").Call("getElementById", ROOT_ID).Set("innerHTML", componentHTML.String())
		js.Global().Get("window").Get("history").Call("pushState", "Object", targetPage.Title, targetPage.Url())
		return nil
	})

	cards := js.Global().Get("document").Call("querySelectorAll", "#"+ids.CardContainer)
	for i := 0; i < cards.Length(); i++ {
		card := cards.Index(i)
		card.Call("addEventListener", "click", clickHandler)
	}
}
