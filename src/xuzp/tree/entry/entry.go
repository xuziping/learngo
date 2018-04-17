package main

import (
	"xuzp/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder () {
	if myNode == nil || myNode.node == nil{
		return
	}
	lNode := myTreeNode{myNode.node.Left}
	lNode.postOrder()
	rNode := myTreeNode{myNode.node.Right}
	rNode.postOrder()
	myNode.node.Print()

}


func main() {

	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}

	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()

	root.Right.Left.SetValue(4)

	root.SetValue(100)

	root.Traverse()

	proot := &root
	proot.Print()
	proot.SetValue(50)
	proot.Print()


	fmt.Println()
	rootNode := myTreeNode{&root}
	rootNode.postOrder()
	//nodes := []treeNode {
	//	{value: 3},
	//	{},
	//	{},
	//	{6, nil, &root},
	//}
	//
	//fmt.Println(nodes)


	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node){
		nodeCount++
	})
	fmt.Println("Node count: ", nodeCount)

}

