package functions

import (
	"ClassTree"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"syscall"
)

func InputDataToTree(
	sortNumber int, isHead bool,
	isInputFromFile, isOutputToFile, isReverse bool,
	nameInputFile, nameOutputFile string) *ClassTree.TopBinaryTree {
	var (
		inputStr      string
		n             int
		oneLine       []string
		err           error
		inputHeadNode ClassTree.TopBinaryTree
		inputTree     = new(ClassTree.TopBinaryTree)
		scanner       *bufio.Scanner
		// out           *os.File
	)

	//############
	// check flags
	if isInputFromFile { //input from file
		os.Stdin, err = os.Open(nameInputFile)

		defer func() { // close input file and set standard os.Stdin
			err = os.Stdin.Close()
			if err != nil {
				log.Fatal(err)
			}
			os.Stdin = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		}()

		if err != nil {
			log.Fatalln(err.Error())
		}
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		//input from console
		os.Stdout.Write([]byte("Enter data:\n"))
		scanner = bufio.NewScanner(os.Stdin)
	}
	if isOutputToFile {
		os.Stdout, err = os.Create(nameOutputFile)

		defer func() { // close output file and set standard os.Stdout
			os.Stdout.Close()
			os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
			//out = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
		}()

		if err != nil {
			log.Fatalln(err.Error())
		}

	}

	if isHead {

		scanner.Scan()
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		headNode := func() {
			inputHeadNode.BinaryNode = new(ClassTree.BinaryTree)
			initHeadNode := inputHeadNode.InitTree(inputHeadNode.BinaryNode)
			initHeadNode(strings.Split(scanner.Text(), ","), 0)

		}
		headNode()
	}
	//#########################

	inputTree.BinaryNode = new(ClassTree.BinaryTree)
	inputTree.HeadNode = inputTree.BinaryNode
	initBranch := inputTree.InitTree(inputTree.BinaryNode)

	//################
	// inputting data
	for scanner.Scan() {

		inputStr = scanner.Text()
		if scanner.Err() != nil {
			log.Fatalf("Error of input! Err: %v", scanner.Err())
		}

		if inputStr == "" {
			break
		}

		oneLine = strings.Split(inputStr, ",")

		if (n != 0 && n != len(oneLine)) || len(oneLine)-1 < sortNumber {
			log.Fatalln("Error of count values!")
		}
		n = len(oneLine)

		initBranch(oneLine, sortNumber)

	}
	//___________

	//#######################
	// write sorted date to Stdout
	if inputHeadNode.BinaryNode != nil {
		inputHeadNode.WriteOnlyInTree(inputHeadNode.BinaryNode, isReverse, true, os.Stdout)
	}
	inputTree.WriteOnlyInTree(inputTree.BinaryNode, isReverse, false, os.Stdout)
	//_________
	return inputTree

}

func InputFromFile(inputFileName string, sortNumber int) [][]string {
	var (
		inputStr   string
		n          int
		arrayLines [][]string
	)

	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		inputStr = scanner.Text()

		if scanner.Err() != nil {
			log.Fatalf("Error of input! Err: %v", scanner.Err())
		}

		arrayStr := strings.Split(inputStr, ",")

		if (n != 0 && n != len(arrayStr)) || len(arrayStr)-1 < sortNumber {
			log.Fatalln("Error of count values!")
		}
		n = len(arrayStr)

		arrayLines = append(arrayLines, arrayStr)
	}

	return arrayLines
}

func WriteToFile(fileName string, arrayLines [][]string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, line := range arrayLines {
		countElements := len(line)
		for i, el := range line {
			str := el
			if i < (countElements - 1) {
				str += ","
			}
			f.WriteString(str)
		}
		f.WriteString("\n")
	}
}
func WriteToConsole(arrayLines [][]string) {
	for _, line := range arrayLines {
		countElements := len(line)
		for i, el := range line {
			if i < (countElements - 1) {
				fmt.Printf("%s,", el)
			} else {
				fmt.Println(el)
			}

		}

	}
}
func SortArray(isFirstHeader, isReverse bool, sortNumber int, array [][]string) [][]string {

	var sortingArray [][]string
	copyOfArray := make([][]string, len(array))
	copy(copyOfArray, array)

	if isFirstHeader {
		sortingArray = copyOfArray[1:]
	} else {
		sortingArray = copyOfArray
	}

	if !isReverse {
		sort.Slice(sortingArray, func(i, j int) bool {
			return sortingArray[i][sortNumber] < sortingArray[j][sortNumber]
		})
	} else {
		sort.Slice(sortingArray, func(i, j int) bool {
			return sortingArray[i][sortNumber] > sortingArray[j][sortNumber]
		})
	}

	return copyOfArray
}

/*
func SortBinTreeArray(isFirstHeader, isReverse bool, sortNumber int, array [][]string) [][]string {

	var sortingArray [][]string
	copyOfArray := make([][]string, len(array))
	copy(copyOfArray, array)

	if isFirstHeader {
		sortingArray = copyOfArray[1:]
	} else {
		sortingArray = copyOfArray
	}

	dataTree := new(ClassTree.TopBinaryTree)
	dataTree.HeadNode = new(ClassTree.BinaryTree)
	dataTree.BinaryNode = dataTree.HeadNode

	addNodeToTree := dataTree.InitTree(dataTree.HeadNode)

	for _, el := range sortingArray {
		addNodeToTree(el, sortNumber)
	}

	arrayStr := dataTree.WriteSorted(dataTree.HeadNode, isReverse, [][]string{})

	if isFirstHeader {
		for i := 1; i < len(copyOfArray); i++ {
			copyOfArray[i] = arrayStr[i-1]
		}
	} else {
		copyOfArray = arrayStr
	}

	return copyOfArray
}
*/
