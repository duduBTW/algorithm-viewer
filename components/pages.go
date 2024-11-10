package components

import "github.com/a-h/templ"

type Controller interface {
	Init()
}

type Page struct {
	Id        string
	Title     string
	Component func() templ.Component
}

func (page Page) Url() string {
	return "/algorithm/" + page.Id
}

var AllAlgorithmsPages = []Page{
	{Component: TwoSum, Id: "two-sum", Title: "Two sum"},
	{Component: CreateBynaryTree, Id: "create-binary-tree", Title: "Create binary tree"},
	{Component: BalanceBinaryTree, Id: "balance-binary-tree", Title: "Balance binary tree"},
}
