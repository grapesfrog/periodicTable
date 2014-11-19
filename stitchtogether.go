// usage ./filecontents -folder=folder_to_walk -outfile=output_file_to_create

/*
This  expects a folder full of valid JSON snippets in individual files .
It will knit the  individual files together as  single valid  JSON file that can be uploaded to mnogodb , bigquery etc
 NOTE: probably should not have used walk as that is recursive but it's so neat a way to work I could also have used a boolean to see if last file being processed
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Folder_to_search string
var File_to_create string
var Record_val string

func main() {

	// flag( name, default value, description)
	flag.StringVar(&Folder_to_search, "folder", "/tmp/elements/", "File to read")
	flag.StringVar(&File_to_create, "outfile", "/tmp/elements.json", "json File to write contents to")
	flag.Parse()

	// flag to check if last line
	last_line_check := 0

	// count of number of files in folder
	Count_of_files, _ := ioutil.ReadDir(Folder_to_search)

	// fmt.Println(len(Count_of_files))
	// fmt.Printf("Number of files : %d \n", len(Count_of_files))

	createNewFile()
	total := 0
	filepath.Walk(Folder_to_search, func(path string, info os.FileInfo, err error) error {
		total++

		fmt.Println(path)
		b, err := ioutil.ReadFile(path)
		Record_val = string(b)

		// total includes count of directory
		if total == len(Count_of_files)+1 {
			last_line_check = 1
		}

		if total > 1 {
			appendToFile(Record_val, last_line_check)
		}
		return nil
	})
	fmt.Printf("Total articles: %d \n", len(Count_of_files))
	endOfFile()
}

func createNewFile() {
	// overwrites any existing file
	outputFile, err := os.Create(File_to_create)
	if err != nil {
		// handle error here
		fmt.Println(err)
		return
	}
	// automatically call Close() at the end of current method
	defer outputFile.Close()
	outputFile.WriteString("[ \n")

}

func endOfFile() {
	// file must already exist
	outputFile, err := os.OpenFile(File_to_create, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	if _, err = outputFile.WriteString("] \n"); err != nil {
		panic(err)
	}
}

func appendToFile(Record_val string, last_line_check int) {
	// file must already exist, if last line
	outputFile, err := os.OpenFile(File_to_create, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()
	// If last file do not write  ","  as part of record to file
	if last_line_check != 1 {

		if _, err = outputFile.WriteString(Record_val + ", \n"); err != nil {
			panic(err)
		}
	} else {
		if _, err = outputFile.WriteString(Record_val + "\n"); err != nil {
			panic(err)
		}
	}
}
