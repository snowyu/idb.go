package iDB

import (
    "testing"
    . "github.com/snowyu/prettytest"
)

func TestUtilsTest(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	Run(
    //RunWithFormatter(
		t,
        //&BDDFormatter{Description: "Utils Testing"},
		new(UtilsTestSuite),
	)
}
