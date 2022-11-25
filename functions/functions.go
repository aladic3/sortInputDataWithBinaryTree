package functions

import (
	"ClassTree"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
)

func ReadDir(path string) chan string {
	fnames := make(chan string)
	go func() {
		defer close(fnames)
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}

			if !info.IsDir() {
				fnames <- info.Name()
			}
			return nil
		})
		if err != nil {
			log.Fatalln(err)
		}
	}()
	return fnames
}

func CheckFlags(isHead, isInputFromFile, isOutputToFile, isInputWithTree bool,
	nameInputFile, nameOutputFile string,
	inputHeadNode, inputTree *ClassTree.TopBinaryTree) (func([]string, int), bool, *bufio.Scanner) {

	var err error
	var returnFunc func([]string, int)
	var result bool
	var scanner *bufio.Scanner

	//############
	// check flags
	if isInputFromFile { //input from file
		findCSV := strings.Split(nameInputFile, ".")
		if findCSV[len(findCSV)-1] != "csv" &&
			findCSV[len(findCSV)-1] != "CSV" {
			log.Fatalf("Input file name must be .csv!")
		}

		os.Stdin, err = os.Open(nameInputFile)
		if err != nil {
			log.Fatalln(err.Error())
		}

	} else {
		//input from console
		_, err = os.Stdout.Write([]byte("Enter data:\n"))
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	scanner = bufio.NewScanner(os.Stdin)

	if isOutputToFile {
		findCSV := strings.Split(nameOutputFile, ".")
		if findCSV[len(findCSV)-1] != "csv" &&
			findCSV[len(findCSV)-1] != "CSV" {
			log.Fatalf("Output file name must be .csv!")
		}

		os.Stdout, err = os.Create(nameOutputFile)
		if err != nil {
			log.Fatalln(err.Error())
		}

	}

	if isHead && isInputWithTree {

		scanner.Scan()
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		func() {
			initHeadNode := inputHeadNode.InitTree()
			initHeadNode(strings.Split(scanner.Text(), ","), 0)

		}()

	}
	//#########################
	if isInputWithTree {
		initTree := inputTree.InitTree()
		returnFunc = initTree
		result = true

	}

	return returnFunc, result, scanner
}

func InputtingData(
	sortNumber int,
	isHead, isInputFromFile, isOutputToFile, isReverse, isInputWithTree bool,
	nameInputFile, nameOutputFile string) *ClassTree.TopBinaryTree {
	var (
		inputStr      string
		n             int
		oneLine       []string
		arrayLines    [][]string
		inputHeadNode = new(ClassTree.TopBinaryTree)
		inputTree     = new(ClassTree.TopBinaryTree)
		err           error
	)

	initBranch, isBranchInput, scanner := CheckFlags(isHead, isInputFromFile, isOutputToFile, isInputWithTree,
		nameInputFile, nameOutputFile,
		inputHeadNode, inputTree)

	defer func() { // close input file and set standard os.Stdin
		err = os.Stdin.Close()
		if err != nil {
			log.Fatal(err)
		}
		os.Stdin = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	}()
	defer func() { // close output file and set standard os.Stdout
		err = os.Stdout.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
		os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")

	}()

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

		// inputting data
		if isBranchInput {
			initBranch(oneLine, sortNumber)
		} else {
			arrayLines = append(arrayLines, oneLine)
		}

	}
	//___________

	//#######################
	// write sorted date to Stdout
	if isBranchInput {
		if inputHeadNode.BinaryNode != nil {
			inputHeadNode.WriteOnlyInTree(inputHeadNode.BinaryNode, isReverse, true, os.Stdout)
		}
		inputTree.WriteOnlyInTree(inputTree.BinaryNode, isReverse, false, os.Stdout)

	}
	if !isBranchInput {
		arrayLines = SortArray(isHead, isReverse, sortNumber, arrayLines)
		for i := 0; i < len(arrayLines); i++ {
			_, err = os.Stdout.Write([]byte(inputTree.BinaryNode.StringifyData(arrayLines[i])))
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
		inputTree = nil
	}
	return inputTree
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
