package algorithms

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (node *Node) Add(newNode *Node) {
	if node.Val > newNode.Val {
		if node.Left != nil {
			node.Left.Add(newNode)
		} else {
			node.Left = newNode
		}
		return
	}

	if node.Val < newNode.Val {
		if node.Right != nil {
			node.Right.Add(newNode)
		} else {
			node.Right = newNode
		}
		return
	}
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Add(val int) {
	newNode := Node{Val: val}

	if tree.Root == nil {
		tree.Root = &newNode
		return
	}

	tree.Root.Add(&newNode)
}

func CreateBinaryTree(vals []int) BinaryTree {
	tree := BinaryTree{}
	for _, val := range vals {
		tree.Add(val)
	}

	return tree
}
