package FunctionsForSort

var (
	info *infoOfSorting
)

type LstBuffer struct {
	data []string
	next *LstBuffer
}

type infoOfSorting struct {
	increaseOrLess bool // 1 if increases, 0 -- decreases
	columnOfSort   int
	head           *LstBuffer
}

func NewLstBuffer() *LstBuffer {
	return &LstBuffer{
		data: nil,
		next: nil,
	}
}

func newNode(data []string) *LstBuffer {
	return &LstBuffer{
		data: data,
		next: nil,
	}
}

func (*LstBuffer) recursiveAdd(tempNode, createdNode *LstBuffer, parameters *infoOfSorting) (result *LstBuffer) {
	result = tempNode.next
	if tempNode != nil {
		if tempNode.data[parameters.columnOfSort] < createdNode.data[parameters.columnOfSort] {
			tempNode.next = createdNode.recursiveAdd(tempNode.next, createdNode, parameters)
		} else {
			if tempNode == parameters.head {
				parameters.head = createdNode
			}
			createdNode.next = tempNode
			result = createdNode

		}
	} else {
		result = createdNode
	}

	return
}

func (sortingLst *LstBuffer) addDataToLstIncrease(data []string, parameters *infoOfSorting) {
	if parameters.head.data == nil {
		parameters.head.data = data
		return
	}

	var (
		createdNode = newNode(data)
		tempNode    = parameters.head
	)

	sortingLst.recursiveAdd(tempNode, createdNode, parameters)

}

func (sortingLst *LstBuffer) SortAndAddData(input <-chan []string, columnOfSort int, increaseOrLess bool) {
	if info == nil {
		info = &infoOfSorting{
			increaseOrLess: increaseOrLess,
			columnOfSort:   columnOfSort,
			head:           sortingLst,
		}
	}
	parameters := info

	for data := range input {
		sortingLst.addDataToLstIncrease(data, parameters)
	}

	if parameters.increaseOrLess {

	} else {

	}

}
