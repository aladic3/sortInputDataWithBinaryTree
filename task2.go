package main

import (
	"flag"
	"log"
	"os"
	"sortInputDataWithBinaryTree/FunctionsForSort"
)

var fileNames chan string
var nChanFiles = 1

func initLogs() func() {
	var (
		outLogFile *os.File
		err        error
	)
	// init logs
	if outLogFile, err = os.OpenFile("logs/logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(outLogFile)
	log.Println("Starting program...")
	return func() {
		outLogFile.Close()
	}

}

func main() {
	closeLogs := initLogs()
	defer closeLogs()

	var (
		inputFileName    = flag.String("i", "", "Use a file with the name file-name as an input")
		outputFileName   = flag.String("o", "", "Use a file with the name file-name as an output")
		isFirstHeader    = flag.Bool("h", flag.Parsed(), "The first line is a header that must be ignored during sorting but included in the output.")
		isReverse        = flag.Bool("r", flag.Parsed(), "Sort input lines in reverse order.")
		sortNumber       = flag.Int("f", 0, "Sort input lines by value number N.")
		sortingAlgorithm = flag.Int("a", 2, "the ability to use a second algorithm for sorting - "+
			"the Tree Sort algorithm. Accordingly, add one more option -a with possible values 1 or 2, which chooses "+
			"currently implemented algorithm or Tree Sort algorithm to use. By default, the application uses the currently implemented algorithm")
		dirName = flag.String("d", "Path with files", "The application has additional option -d dir-name that specifies a directory where it must read"+
			"input files from. All files in the directory must have the same format. The output stays the same, it "+
			"is a one file or a standard output with sorted content from all input files. ")
	)
	flag.Parse()

	var (
		isInputFromFile = (*inputFileName != "") || (*dirName != "")
		isOutputToFile  = *outputFileName != ""
		isInputWithTree = *sortingAlgorithm == 2
	)

	if *sortingAlgorithm > 2 || *sortingAlgorithm < 1 {
		log.Fatal("Bad flag 'a'")
	}

	if *dirName != "" && *inputFileName != "" {
		log.Fatal("Don't use 2 options -i and -d!!!")
	}

	fileNames = FunctionsForSort.ReadDir(*dirName, *inputFileName, &nChanFiles)

	FunctionsForSort.InputtingAndSortingData(*sortNumber,
		*isFirstHeader, isInputFromFile, isOutputToFile, *isReverse, isInputWithTree,
		*outputFileName,
		fileNames, FunctionsForSort.Handler(), nChanFiles)

}
