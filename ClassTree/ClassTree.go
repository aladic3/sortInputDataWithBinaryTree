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

func (Tree *BinaryTree) StringifyData(data []string) string {
	var line string
	var countElements int
	var sfyData []string
	if data == nil {
		countElements = len(Tree.data)
		sfyData = Tree.data
	} else {
		countElements = len(data)
		sfyData = data
	}

	for i := 0; i < countElements; i++ {
		line += sfyData[i]
		if i < (countElements - 1) {
			line += ","
		}
	}
	line += "\n"

	return line
}

func (ThisTree *TopBinaryTree) InitTree() func(
	data []string, column int) {
	var isFirstRead = true
	ThisTree.HeadNode = new(BinaryTree)
	ThisTree.BinaryNode = ThisTree.HeadNode
	rootNode := ThisTree.HeadNode

	return func(data []string, column int) {
		CurrentNode := rootNode

		if isFirstRead {
			CurrentNode.data = data
			isFirstRead = false
			return
		}

		for {
			var ok bool
			if CurrentNode.data[column] <= data[column] {
				if CurrentNode.right == nil {
					CurrentNode.right, ok = CurrentNode.setCurrentlyNode(
						CurrentNode.right, data)
					if ok {
						return
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
						return
					}
				} else {
					CurrentNode = CurrentNode.left
				}

			}
		}

	}

}

func (Tree *TopBinaryTree) WriteOnlyInTree(Node *BinaryTree,
	isReverse bool, isHead bool,
	output *os.File) {
	var countByteWrite int
	var err error
	_ = countByteWrite

	if isHead {
		countByteWrite, err = output.Write([]byte(Node.StringifyData(nil)))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if Node != nil && !isReverse {
		Tree.WriteOnlyInTree(Node.left, isReverse, false, output)
		countByteWrite, err = output.Write([]byte(Node.StringifyData(nil)))
		if err != nil {
			log.Fatal(err)
		}
		Tree.WriteOnlyInTree(Node.right, isReverse, false, output)
	}
	if Node != nil && isReverse {
		Tree.WriteOnlyInTree(Node.right, isReverse, false, output)
		countByteWrite, err = output.Write([]byte(Node.StringifyData(nil)))
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
