package main

import (
	"ClassTree"
	"flag"
	"functions"
)

func main() {

	var (
		arrayTrees []*ClassTree.TopBinaryTree

		inputFileName   = flag.String("i", "", "Use a file with the name file-name as an input")
		outputFileName  = flag.String("o", "", "Use a file with the name file-name as an output")
		isFirstHeader   = flag.Bool("h", flag.Parsed(), "The first line is a header that must be ignored during sorting but included in the output.")
		isReverse       = flag.Bool("r", flag.Parsed(), "Sort input lines in reverse order.")
		sortNumber      = flag.Int("f", 0, "Sort input lines by value number N.")
		sortingAlgoritm = flag.Int("a", 2, "the ability to use a second algorithm for sorting - "+
			"the Tree Sort algorithm. Accordingly, add one more option -a with possible values 1 or 2, which chooses "+
			"currently implemented algorithm or Tree Sort algorithm to use. By default, the application uses the currently implemented algorithm")
	)
	flag.Parse()

	if *sortingAlgoritm == 2 {
		lambda := functions.InputDataToTree
		oneTree := lambda(*sortNumber,
			*isFirstHeader, *inputFileName != "", *outputFileName != "", *isReverse,
			*inputFileName, *outputFileName)
		arrayTrees = append(arrayTrees, oneTree)

		return
	}
	//if *inputFileName == "" {
	//	arrayLines = functions.InputFromConsole(*sortNumber, true)
	//} else {
	//	arrayLines = functions.InputFromFile(*inputFileName, *sortNumber)
	//}

	//switch *sortingAlgoritm {
	//case 1:
	//	arrayLines = functions.SortArray(*isFirstHeader, *isReverse, *sortNumber, arrayLines)
	//case 2:
	//	arrayLines = functions.SortBinTreeArray(*isFirstHeader, *isReverse, *sortNumber, arrayLines)
	//default:
	//	arrayLines = functions.SortArray(*isFirstHeader, *isReverse, *sortNumber, arrayLines)
	//}
	//
	//if *outputFileName == "" {
	//	functions.WriteToConsole(arrayLines)
	//} else {
	//	functions.WriteToFile(*outputFileName, arrayLines)
	//}
}
