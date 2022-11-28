package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const countLines = 1000000

func main() {
	ch := make(chan string)
	fileName := "GeneratedFile.csv"

	generator(ch, countLines)

	writeToFile(fileName, ch)

	fmt.Println(getSizeFile(fileName))

}

func writeToFile(fileName string, ch chan string) {
	file, _ := os.Create(fileName)
	defer file.Close()

	for line := range ch {
		file.WriteString(line + "\n")
	}
}

func getSizeFile(fileName string) int64 {
	f, er := os.Open(fileName)
	if er != nil {
		log.Fatal(er)
	}
	fStat, _ := f.Stat()
	return fStat.Size()

}

func generator(ch chan string, countLines int) {

	go func() {
		defer close(ch)

		fields := make([]string, 3)
		for i := 0; i < countLines; i++ {
			fields[0] = fmt.Sprintf("f%d", i)
			fields[1] = fmt.Sprintf("s%d", i)
			fields[2] = fmt.Sprintf("t%d", i)
			ch <- strings.Join(fields, ",")
		}

	}()
}
