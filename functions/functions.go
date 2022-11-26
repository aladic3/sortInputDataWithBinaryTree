package functions

import (
	"ClassTree"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
)

func ReadDir(path string) chan string {
	fnames := make(chan string)
	if path == "" {
		go func() {
			defer close(fnames)
			fnames <- path
		}()
		return fnames
	}
	go func(ch chan string) {
		defer close(ch)
		var pathToFile string
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}

			if info.IsDir() {
				pathToFile = pathToFile + info.Name() + "\\"
			}
			if !info.IsDir() {
				ch <- path
			}
			return nil
		})
	}(fnames)
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

		scanner = bufio.NewScanner(inputFile)
	} else {
		//input from console
		_, err = os.Stdout.Write([]byte("Enter data:\n"))
		if err != nil {
			log.Fatalln(err.Error())
		}
		scanner = bufio.NewScanner(os.Stdin)
	}

	if isHead {

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
	nameOutputFile string,
	fnames <-chan string,
	nChan int) *ClassTree.TopBinaryTree {
	var (
		n             int
		arrayLines    [][]string
		inputHeadNode = new(ClassTree.TopBinaryTree)
		inputTree     = new(ClassTree.TopBinaryTree)
		err           error
		initTree      = func([]string, int) {}
		allLines      = make(chan string)
	)
	// init input func is input with tree algorithm
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

	// fan-out + processing
	lines := make([]chan string, nChan)
	for i := 0; i < nChan; i++ {
		lines[i] = make(chan string)
		go func(line chan string) {
			defer close(line)

			for fname := range fnames {
				//create scanner
				scanner, inputFile := CheckFlags(isHead, isInputFromFile, isInputWithTree,
					fname,
					inputHeadNode)

				for scanner.Scan() {

					inputStr := scanner.Text()
					if scanner.Err() != nil {
						log.Fatalf("Error of input! Err: %v", scanner.Err())
					}

					if inputStr == "" {
						break
					}

					oneLine := strings.Split(inputStr, ",")

					if (n != 0 && n != len(oneLine)) || len(oneLine)-1 < sortNumber {
						log.Fatalln("Error of count values!")
					}
					n = len(oneLine)

					// inputting data to chan
					line <- inputStr

				}

				// close input file
				if isInputFromFile {
					defer func() {
						err = inputFile.Close()
						if err != nil {
							log.Fatal(err)
						}
					}()
				}
			}
		}(lines[i])
	}

	// fan-in
	go func() {
		defer close(allLines)
		wg := &sync.WaitGroup{}

		for i := range lines {
			wg.Add(1)
			go func(ch chan string) {
				defer wg.Done()
				for line := range ch {
					allLines <- line
				}
			}(lines[i])
		}

		wg.Wait()
	}()

	// inputting data and sorting
	if isInputWithTree {
		for content := range allLines {
			initTree(strings.Split(content, ","), sortNumber)
		}
	} else {
		arrayLines = SortArray(isReverse, sortNumber, allLines)
	}
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
		if isHead {
			inputHeadNode.WriteOnlyInTree(inputHeadNode.BinaryNode, isReverse, true, os.Stdout)
		}
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

func SortArray(isReverse bool, sortNumber int, content <-chan string) [][]string {
	var sortingArray [][]string
	var buffer = make([][]string, 0, 1000)
	//copyOfArray := make([][]string, len(array))
	//copy(copyOfArray, array)
	/*
		if isFirstHeader {
			sortingArray = copyOfArray[1:]
		} else {
			sortingArray = copyOfArray
		}
	*/

	// read lines -> buffer
	for line := range content {
		buffer = append(buffer, strings.Split(line, ","))
	}

	sortingArray = buffer

	if !isReverse {
		sort.Slice(sortingArray, func(i, j int) bool {
			return sortingArray[i][sortNumber] < sortingArray[j][sortNumber]
		})
	} else {
		sort.Slice(sortingArray, func(i, j int) bool {
			return sortingArray[i][sortNumber] > sortingArray[j][sortNumber]
		})
	}

	return sortingArray
}
