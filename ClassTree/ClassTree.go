package ClassTree

import (
	"log"
	"os"
)

type TopBinaryTree struct {
	BinaryNode *BinaryTree
	HeadNode   *BinaryTree
}
type BinaryTree struct {
	data     []string
	left     *BinaryTree
	right    *BinaryTree
	HeadNode *BinaryTree
}

func (Tree *BinaryTree) StringifyData() string {
	var line string
	countElements := len(Tree.data)
	for i := 0; i < countElements; i++ {
		line += Tree.data[i]
		if i < (countElements - 1) {
			line += ","
		}
	}
	line += "\n"

	return line
}

func (ThisTree *TopBinaryTree) InitTree(
	rootNode *BinaryTree) func(
	data []string, column int) *BinaryTree {
	var isFirstRead = true

	return func(data []string, column int) *BinaryTree {
		CurrentNode := rootNode

		if isFirstRead {
			CurrentNode.data = data
			isFirstRead = false
			return rootNode
		}

		finish := false

		for !finish {
			var ok bool
			if CurrentNode.data[column] <= data[column] {
				if CurrentNode.right == nil {
					CurrentNode.right, ok = CurrentNode.setCurrentlyNode(
						CurrentNode.right, data)
					if ok {
						return rootNode
					}
				} else {
					CurrentNode = CurrentNode.right

				}

			}

			if CurrentNode.data[column] > data[column] {
				if CurrentNode.left == nil {
					CurrentNode.left, ok = CurrentNode.setCurrentlyNode(
						CurrentNode.left, data)
					if ok {
						return rootNode
					}
				} else {
					CurrentNode = CurrentNode.left
				}

			}
		}
		return rootNode
	}

}

func (Tree *TopBinaryTree) WriteOnlyInTree(Node *BinaryTree,
	isReverse bool, isHead bool,
	output *os.File) {
	var countByteWrite int
	var err error
	_ = countByteWrite

	if isHead {
		countByteWrite, err = output.Write([]byte(Node.StringifyData()))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if Node != nil && !isReverse {
		Tree.WriteOnlyInTree(Node.left, isReverse, false, output)
		countByteWrite, err = output.Write([]byte(Node.StringifyData()))
		if err != nil {
			log.Fatal(err)
		}
		Tree.WriteOnlyInTree(Node.right, isReverse, false, output)
	}
	if Node != nil && isReverse {
		Tree.WriteOnlyInTree(Node.right, isReverse, false, output)
		countByteWrite, err = output.Write([]byte(Node.StringifyData()))
		if err != nil {
			log.Fatal(err)
		}
		Tree.WriteOnlyInTree(Node.left, isReverse, false, output)

	}

	if Node == nil {
		return
	}
	return

}

func (ThisTree *BinaryTree) setCurrentlyNode(
	Node *BinaryTree, data []string) (*BinaryTree, bool) {
	var end bool
	if Node == nil {
		Node = new(BinaryTree)
		Node.data = data
		end = true

	}

	return Node, end
}
