//go:build js && wasm
// +build js,wasm

package components

type ClientPage struct {
	Id        string
	Component BaseComponent
}

func (page ClientPage) Url() string {
	return "/algorithm/" + page.Id
}

var ClientAllAlgorithmsPages = []ClientPage{
	// {Component: TwoSum, Id: "two-sum", Title: "Two sum"},
	{Component: CreateBinaryTreeComponent, Id: "create-binary-tree"},
	// {Component: BalanceBinaryTree, Id: "balance-binary-tree", Title: "Balance binary tree"},
}
