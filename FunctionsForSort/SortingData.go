package FunctionsForSort

type LstBuffer struct {
	data           []string
	head           *LstBuffer
	next           *LstBuffer
	increaseOrLess bool
	columnOfSort   int
}

func (sortingLst *LstBuffer) SortingData(input <-chan []string) {
	var temp *LstBuffer
	if sortingLst.head == nil {
		temp = new(LstBuffer)
		sortingLst.head = temp
		temp.head = sortingLst.head
		temp.data = sortingLst.data
	}

	if sortingLst.increaseOrLess { // first less next
		for massInputData := range input {
			_ = massInputData
		}
	} else {

	}
}
