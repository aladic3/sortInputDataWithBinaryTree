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

func CheckFlags(isHead, isInputFromFile, isInputWithTree bool,
	nameInputFile string,
	inputHeadNode *ClassTree.TopBinaryTree) (*bufio.Scanner, *os.File) {

	var err error
	var scanner *bufio.Scanner
	var inputFile *os.File

	//############
	// check flags
	if isInputFromFile { //input from file
		findCSV := strings.Split(nameInputFile, ".")
		if findCSV[len(findCSV)-1] != "csv" &&
			findCSV[len(findCSV)-1] != "CSV" {
			log.Fatalf("Input file name must be .csv!")
		}

		inputFile, err = os.Open(nameInputFile)
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
	scanner = bufio.NewScanner(inputFile)

	if isHead && isInputWithTree {

		scanner.Scan()
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		if inputHeadNode.HeadNode == nil {
			func() {
				initHeadNode := inputHeadNode.InitTree()
				initHeadNode(strings.Split(scanner.Text(), ","), 0)

			}()
		} else {
			scanner.Text()
		}

	}
	//#########################

	return scanner, inputFile
}

func InputtingData(
	sortNumber int,
	isHead, isInputFromFile, isOutputToFile, isReverse, isInputWithTree bool,
	nameInputFile, nameOutputFile string,
	fnames <-chan string) *ClassTree.TopBinaryTree {
	var (
		inputStr      string
		n             int
		oneLine       []string
		arrayLines    [][]string
		inputHeadNode = new(ClassTree.TopBinaryTree)
		inputTree     = new(ClassTree.TopBinaryTree)
		err           error
		initTree      = func([]string, int) {}
	)
	if isInputWithTree {
		initTree = inputTree.InitTree()
	}

	// init stdout output
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
	defer func() { // close output file and set standard os.Stdout
		err = os.Stdout.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
		os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")

	}()

	//################
	// inputting data

	go func() {
		for fname := range fnames {
			scanner, inputFile := CheckFlags(isHead, isInputFromFile, isInputWithTree,
				fname,
				inputHeadNode)

			defer func() { // close input file
				err = inputFile.Close()
				if err != nil {
					log.Fatal(err)
				}
			}()

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
				if isInputWithTree {
					initTree(oneLine, sortNumber)
				} else {
					arrayLines = append(arrayLines, oneLine)
				}

			}

		}
	}()

	//___________

	//#######################
	// write sorted date to Stdout
	if isInputWithTree {
		if inputHeadNode.BinaryNode != nil {
			inputHeadNode.WriteOnlyInTree(inputHeadNode.BinaryNode, isReverse, true, os.Stdout)
		}
		inputTree.WriteOnlyInTree(inputTree.BinaryNode, isReverse, false, os.Stdout)

	}
	if !isInputWithTree {
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
