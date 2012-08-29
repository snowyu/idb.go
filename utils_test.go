package iDB

import (
    //"os"
    "testing"
    . "github.com/remogatto/prettytest"
)


type utilsSuite struct {
  Suite
  testfile string
}

func (suite *utilsSuite) Before() {
    suite.testfile = "mytestfile"
}

func (Suite *utilsSuite) After() {
    os.Remove(Suite.testfile)
}

const testfile  = "mytestfile"

func Test_SetXatrr(t *testing.T){ //test function starts with "Test" and takes a pointer to type testing.T
    err := SetXattr(testfile, "mykey", "myvalue23") 
    if (err != nil || !IsXattrExists(testfile, "mykey")) { //try a unit test on function
        t.Error("SetXattr did not work as expected. Error:", err) // log error if it did not work as expected
    } else {
        t.Log("one test passed.") // log some info if you want
    }
}
func Test_GetXatrr(t *testing.T){
    Test_SetXatrr(t)
    result, err := GetXattr(testfile, "mykey")
    if (result != "myvalue23") { //try a unit test on function
        t.Error("GetXattr did not work as expected. Error:", err) // log error if it did not work as expected
    } else {
        t.Log("one test passed.") // log some info if you want
    }
}

func Test_FileIsExists(t *testing.T){
    result, err := FileIsExists(testfile)
    if !result {
        t.Error("FileIsExists did not work as expected, Error:", err)
    } else {
        t.Log("one test passed")
    }
}
