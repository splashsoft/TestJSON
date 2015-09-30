package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"math"
)


func exercise_12_7() {
	inputFile, _ := os.Open("goprogram.go")
	outputFile, _ := os.OpenFile("goprogramT.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0666)

	defer inputFile.Close()
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	defer outputWriter.Flush() 

	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}

		fmt.Printf("\nI: '%s' - Länge: %d - cap: %d ...\n", string(inputString), len(inputString), cap(inputString))
		// hier hat der slice hinter die eigentlichen Daten gegriffen, hat nur nicht geknallt, weil die cap() größer war
		// soll das so sein?!?
		outputString := "          "
		if( len(inputString) > 2 ) {
			nMaxIdx := int(math.Min( 12, float64(len(inputString)))) 
			fmt.Printf("grabbing inputString[2:%d] ...\n", nMaxIdx)
			outputString = string(inputString[2:nMaxIdx])
			}
		fmt.Printf("O: '%s' - Länge: %d ...\n------------------------------------------------------------------------------------------------------------------\n", outputString, len(outputString))
		
		
		outputString += "\r\n"
		
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		} else {
//			if n < 0 {
//			fmt.Printf("'%s' - %d bytes written ...\n", outputFile.Name(), n )
//			}
		}
	}
	
	fmt.Println("Conversion done")
}

