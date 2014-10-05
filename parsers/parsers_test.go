package parsers

/* Some simple  GO default tests as part of the GO learning experience -
   Ideally you write the tests before coding.  You can stu them out see example
   Creating the tests first actually helps formulate how your code should work
   You then write yor code to pass the tests.
*/

import (
	"fmt"
	"os"
	"testing"
)

// Note: as no return value when creating file there is no if stmt to check the return value as part of test  so not textbook

func Test_CreateNewFile_1(t *testing.T) {
	var Outputfile = "/tmp/outputfromtest.json"
	// tidying up before running file create test don't care if it exsts or not so no need to error
	os.Remove("/tmp/outputfromtest.json")
	CreateNewFile(Outputfile)

	// os.Remove("/tmp/outputfromtest.json") //  <-- This was to validate that test was appropriate uncomment to prove it yourself :-)

	// check that file has been created - if created pass
	if _, err := os.Stat(Outputfile); err == nil {
		fmt.Println(Outputfile, "exist! <-- Pass")
		t.Log("Passes because file created")
	} else {
		t.Error("Not creating output file")

	}

}

func Test_CreateNewFile_2(t *testing.T) {
	var Outputfile = ""
	CreateNewFile(Outputfile)

	// check that file has  NOT been created if NO file name passed
	if _, err := os.Stat(Outputfile); err == nil {

		t.Error("Fails because does not exit when trying to create output file with no name")
	} else {
		fmt.Println("Did not create file with no name  <-- Pass")
		t.Log("It should not continue after tryng to Create output file with no name")

	}

}

func Test_AppendToFile_1(t *testing.T) {
	// check that the outputfile exists and it returns without panicing
	t.Error("This will fail ")

}

func Test_AppendToFile_2(t *testing.T) {
	// check that  string passed to LineValue and it returns without panicing
	t.Error("This will fail ")

}
