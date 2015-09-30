package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("goprogram.go")
	outputFile, _ := os.OpenFile("goprogramT.go", os.O_WRONLY|os.O_CREATE,
		0666)

	defer inputFile.Close()
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)

	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		outputString := string(inputString[2:5]) + "\r\n"
		n, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("'%s' - %d bytes written ...\n", outputFile.Name(), n )
		}
	}
	fmt.Println("Conversion done")
}
