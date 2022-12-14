package FunctionsForSort

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"sortInputDataWithBinaryTree/ClassTree"
	"strings"
	"sync"
	"syscall"
)

// ListSorter yet not working
type ListSorter interface {
	SortAndAddData(input <-chan []string)
}

func Handler() chan struct{} {
	interrupt := make(chan struct{})
	sigs := make(chan os.Signal)
	go func() {
		defer close(interrupt)
		signal.Notify(sigs, syscall.SIGINT)
		<-sigs
	}()
	return interrupt
}

func ReadDir(path string, inputFileName string, nChanFiles *int) chan string {
	fNames := make(chan string)
	if path == "" && inputFileName == "" {
		go func() {
			defer close(fNames)
			fNames <- path
		}()
		return fNames
	}

	if path == "" && inputFileName != "" {
		path = inputFileName
	}

	// calculate count files and set: nChanFiles == count files
	count := 0
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if !info.IsDir() {
			count++
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	*nChanFiles = count

	go func(ch chan string) {
		defer close(ch)

		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}

			if !info.IsDir() {
				ch <- path
			}
			return nil
		})
		if err != nil {
			log.Fatalln(err)
		}
	}(fNames)
	return fNames
}

func CheckSomeFlagsAndSetScanner(isHead, isInputFromFile bool,
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

func InitOutputStdout(isOutputToFile bool, nameOutputFile string) func() {
	var err error
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
	return func() { // close output file and set standard os.Stdout
		err = os.Stdout.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
		os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")

	}
}

func ReadData(scanner *bufio.Scanner,
	sortNumber int,
	line chan []string, interrupt <-chan struct{}) {
	var n int

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
		select {
		case line <- oneLine:
		case <-interrupt:
			break
		}

	}
}

func InputtingAndSortingData(
	sortNumber int,
	isHead, isInputFromFile, isOutputToFile, isReverse, isInputWithTree bool,
	nameOutputFile string,
	fNames <-chan string, interrupt <-chan struct{},
	nChan int) *ClassTree.TopBinaryTree {
	var (
		arrayLines    [][]string
		inputHeadNode = new(ClassTree.TopBinaryTree)
		inputTree     = new(ClassTree.TopBinaryTree)
		err           error
		allLines      = make(chan []string)
		done          chan struct{}
	)

	// init stdout output and close
	closeOutputFile := InitOutputStdout(isOutputToFile, nameOutputFile)
	defer closeOutputFile()

	//################
	// inputting data

	// fan-out + processing
	lines := make([]chan []string, nChan)
	for i := 0; i < nChan; i++ {
		lines[i] = make(chan []string)
		go func(line chan []string) {
			defer close(line)

			for fName := range fNames {
				//create scanner
				scanner, inputFile := CheckSomeFlagsAndSetScanner(isHead, isInputFromFile,
					fName,
					inputHeadNode)

				ReadData(scanner, sortNumber, line, interrupt)

				// close input file
				func() {
					if isInputFromFile {
						func() {
							err = inputFile.Close()
							if err != nil {
								log.Fatal(err)
							}
						}()
					}
				}()

			}
		}(lines[i])
	}

	// fan-in
	go func() {
		defer close(allLines)
		wg := &sync.WaitGroup{}

		for i := range lines {
			wg.Add(1)
			go func(ch chan []string) {
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
		done = make(chan struct{})
		initTree := inputTree.InitTree()
		for i := 0; i < nChan; i++ {
			go func() {
				for content := range allLines {
					initTree(content, sortNumber)
				}

				select {
				case <-done: // if close done
				default: // if close allLines
					close(done)
				}

			}()

		}
	} else {
		arrayLines = SortArray(isReverse, sortNumber, allLines)
	}
	//___________

	//#######################
	// write sorted date to Stdout
	if isInputWithTree {
		<-done
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

func SortArray(isReverse bool, sortNumber int, content <-chan []string) [][]string {
	var sortingArray [][]string
	var buffer = make([][]string, 0, 1000)

	// read lines -> buffer
	for line := range content {
		buffer = append(buffer, line)
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
