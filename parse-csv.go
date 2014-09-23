//Author: Grace
/*  why no struts?  well idea is eventually this will not require you to hard code  details for the key value names  and be a more generic bit of code
My current thoughts this will be easier to do if  I grab the header row as a string and break it up that way I may change my mind though on approach. Suggestions welcome
*/

// usage ./parse-csv -ouput=name-of-generated-output-file -input=name-of-csv-file-to-read

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

//Number of elements in the periodic table
const count_of_elements int = 118

// keyname string that wil hold the name  of the key value pair as it builds up a record
var keyname string

var LineValue string
var outputfile string
var inputfile string

func main() {

	// flag( name, default value, description)
	flag.StringVar(&inputfile, "input", "PeriodicTableDataSet.csv", "File to read")
	flag.StringVar(&outputfile, "output", "/tmp/output.json", "JSON file created")
	flag.Parse()

	// set file values <- yeah I now ugly

	inputfile = inputfile
	outputfile = outputfile

	createNewFile()

	// TODO: at some point figure out easiest way to extract header line so can make a generic function for any csv file

	// mykeynames array to hold header row which constitutes the names that are used to create the keys
	mykeynames := []string{"Z", "Symbol", "Name", "A", "N", "Period", "IUPAC_Group", "Old_IUPAC", "CAS_Group", "Category", "Standard_Atomic_Weight", "Last _Digit", "Ionisation_eV", "Normal_State", "Density_Kg_m3_20°C", "Melting_point_°C", "Melting_point_K", "Boiling_point_°C", "Boiling_point_K", "Year_of_discovery", "Discoverer"}

	csvfile, err := os.Open(inputfile)

	if err != nil {
		fmt.Println(err)
		return
	}
	// automatically call Close() at the end of current method
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	// automaticaly defaults to ',' as separator

	LineCount := 0

	for {
		record, err := reader.Read()

		if err == io.EOF {
			/*LineValue = "End of file"
			appendToFile(LineValue)*/
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		//
		keylistSlice := make([]string, len(record))

		// setting up key value pairs
		if LineCount == 0 {

			fmt.Println("[")
			fmt.Println("{")

			copy(keylistSlice, record)

			fmt.Println("Key names :", keylistSlice)

			for i := 0; i < len(keylistSlice); i++ {
				//			fmt.Println(" ", record[i])
				mykeys := strings.TrimSpace(keylistSlice[i])
				fmt.Println("\"" + mykeys + "\"")

			}
		}

		if LineCount != 0 {
			fmt.Println("{")
			LineValue = "{\n"
			appendToFile(LineValue)
			for i := 0; i < len(record); i++ {
				// t the  unimaginative vatriable name that holds the "value" part of the key value pair as it builds up the record
				t := strings.TrimSpace(record[i])
				keyname := strings.TrimSpace(mykeynames[i])
				if t != "" {
					// constructing key value pairing
					fmt.Println("\"" + keyname + "\": " + "\"" + t + "\"")
					LineValue = "\"" + keyname + "\": " + "\"" + t + "\""
					appendToFile(LineValue)
				} else {
					fmt.Println("\"" + keyname + "\": " + "Null" + "\"")
					LineValue = "\"" + keyname + "\": " + "\"Null" + "\""
					appendToFile(LineValue)
				}
				// formating so valid json-  checking if key value , end of record and also if last record

				if (i == len(record)-1) && (LineCount != count_of_elements) {
					fmt.Println("\n },\n")
					LineValue = "\n },\n"
					appendToFile(LineValue)
				} else if (i == len(record)-1) && (LineCount == count_of_elements) {
					fmt.Println("\n }\n")
					LineValue = "\n }\n"
					appendToFile(LineValue)
				} else {
					LineValue = "," + "\n"
					appendToFile(LineValue)
				}

			}

		}

		LineCount += 1
	}
	// closing "]"
	fmt.Println("]")
	LineValue = "] "
	appendToFile(LineValue)

}

func createNewFile() {
	// overwrites any existing file
	outputFile, err := os.Create(outputfile)
	if err != nil {
		// handle error here
		fmt.Println(err)
		return
	}
	// automatically call Close() at the end of current method
	defer outputFile.Close()
	outputFile.WriteString("[ \n")

}

func appendToFile(lineValue string) {
	// file must already exist
	outputFile, err := os.OpenFile(outputfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	if _, err = outputFile.WriteString(LineValue); err != nil {
		panic(err)
	}
}
