package main

import (
	"flag"
	"functions"
	"log"
)

func main() {

	var (
		inputFileName    = flag.String("i", "", "Use a file with the name file-name as an input")
		outputFileName   = flag.String("o", "", "Use a file with the name file-name as an output")
		isFirstHeader    = flag.Bool("h", flag.Parsed(), "The first line is a header that must be ignored during sorting but included in the output.")
		isReverse        = flag.Bool("r", flag.Parsed(), "Sort input lines in reverse order.")
		sortNumber       = flag.Int("f", 0, "Sort input lines by value number N.")
		sortingAlgorithm = flag.Int("a", 1, "the ability to use a second algorithm for sorting - "+
			"the Tree Sort algorithm. Accordingly, add one more option -a with possible values 1 or 2, which chooses "+
			"currently implemented algorithm or Tree Sort algorithm to use. By default, the application uses the currently implemented algorithm")
	)
	flag.Parse()

	if *sortingAlgorithm > 2 || *sortingAlgorithm < 1 {
		log.Fatal("Bad flag 'a'")
	}

	functions.InputtingData(*sortNumber,
		*isFirstHeader, *inputFileName != "", *outputFileName != "", *isReverse, *sortingAlgorithm == 2,
		*inputFileName, *outputFileName)

	return

}
