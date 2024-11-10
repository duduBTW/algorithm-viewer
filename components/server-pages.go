package components

import "github.com/a-h/templ"

type ServerPage struct {
	Id        string
	Title     string
	Component func() templ.Component
}

func (page ServerPage) Url() string {
	return "/algorithm/" + page.Id
}

var ServerAllAlgorithmsPages = []ServerPage{
	// {Component: TwoSum, Id: "two-sum", Title: "Two sum"},
	{Component: CreateBynaryTree, Id: "create-binary-tree", Title: "Create binary tree"},
	// {Component: BalanceBinaryTree, Id: "balance-binary-tree", Title: "Balance binary tree"},
}
