package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var s []int
	return pickVal(root, s)
}

func pickVal(n *Node, s []int) []int {
	if n == nil {
		return nil
	}
	s = append(s, n.Val)
	for _, c := range n.Children {
		s = pickVal(c, s)
	}
	return s
}
