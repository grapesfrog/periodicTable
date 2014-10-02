package parsers

import (
	"fmt"
	"os"
)

func CreateNewFile(Outputfile string) {
	// overwrites any existing file
	outputFile, err := os.Create(Outputfile)
	if err != nil {
		// handle error here
		fmt.Println(err)
		return
	}
	// automatically call Close() at the end of current method
	defer outputFile.Close()
	outputFile.WriteString("[ \n")

}

func AppendToFile(LineValue string, Outputfile string) {
	// file must already exist
	outputFile, err := os.OpenFile(Outputfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	if _, err = outputFile.WriteString(LineValue); err != nil {
		panic(err)
	}
}
