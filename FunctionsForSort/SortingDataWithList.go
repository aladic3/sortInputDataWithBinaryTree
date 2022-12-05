package FunctionsForSort

type LstBuffer struct {
	data           []string
	head           *LstBuffer
	next           *LstBuffer
	increaseOrLess bool // 1 if increases, 0 -- decreases
	columnOfSort   int
}

func (node *LstBuffer) NewLstBuffer(columnOfSort int, increaseOrLess bool, startData <-chan []string) *LstBuffer {
	var list *LstBuffer

	if node == nil {
		list = &LstBuffer{
			data:           nil,
			head:           nil,
			next:           nil,
			increaseOrLess: increaseOrLess, // 1 if increases, 0 -- decreases
			columnOfSort:   columnOfSort,
		}
		list.head = list
	} else {
		list = &LstBuffer{
			data:           nil,
			head:           node.head,
			next:           node.head.next,
			increaseOrLess: increaseOrLess, // 1 if increases, 0 -- decreases
			columnOfSort:   columnOfSort,
		}
	}

	go func() { list.data = <-startData }()

	return list
}

func (sortingLst *LstBuffer) setHeadNode() {
	sortingLst = sortingLst.head
}

func (sortingLst *LstBuffer) addDataToLstIncrease(data []string) {
	var columnOfSort = sortingLst.columnOfSort
	var chanForAdd = make(chan []string)
	var newNode = LstBuffer.NewLstBuffer()

	defer close(chanForAdd)
	defer sortingLst.setHeadNode()

	if sortingLst.data[columnOfSort] < data[columnOfSort] {

	}
}

func (sortingLst *LstBuffer) SortAndAddData(input <-chan []string) {
	var tempNode = sortingLst

	if sortingLst.increaseOrLess {
		for data := range input {
			tempNode.addDataToLstIncrease(data)

		}

	} else {
		for data := range input {

			tempNode = sortingLst.head

			tempNode = sortingLst.head
		}
	}

}
