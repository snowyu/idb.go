package iDB

import (
    "os"
    //"testing"
    "io/ioutil"
    . "github.com/snowyu/prettytest"
)


type UtilsTestSuite struct {
  Suite
  testfile string
}

func (suite *UtilsTestSuite) Before() {
    suite.testfile = "mytestfile"
    result, _ := FileIsExists(suite.testfile)
    if !result {
        ioutil.WriteFile(suite.testfile, nil, 0600)
    }
    suite.Path(suite.testfile)
}

func (suite *UtilsTestSuite) After() {
    os.Remove(suite.testfile)
}

func (suite *UtilsTestSuite) doTestSetXattr(aKey, aValue string){
    err := SetXattr(suite.testfile, aKey, aValue)
    suite.Nil(err)
    suite.True(IsXattrExists(suite.testfile, aKey))
    suite.False(IsXattrExists(suite.testfile, aKey+".myNotExistsTheXKey"))
}

func (suite *UtilsTestSuite) TestSetXattr(){
    suite.doTestSetXattr("mykey", "myvalue23")
    suite.doTestSetXattr(".dsd.sd", "ajsae@3if34$/dsd")
}

func InStrings(aStr string, aList []string) bool {
  result := false
  for i := range aList {
      if aList[i] == aStr {
          result = true
          break
      }
  }
  return result
}

func (suite *UtilsTestSuite) TestListXattr(){
    pairs := map[string]string {
        "mykey": "myvalue23",
        "mykey2":"ajsae@3if34$/dsd",
    }
    for k,v := range pairs {
      suite.doTestSetXattr(k, v)
    }
    names, err := ListXattr(suite.testfile)
    suite.Nil(err)
    for k := range pairs {
        suite.True(InStrings(k, names))
    }
}

func (suite *UtilsTestSuite) TestGetXatrr(){
    //suite.TestSetXattr()
    suite.doTestSetXattr("mykey", "myvalue23")
    pairs := map[string]string {
        "mykey": "myvalue23",
        "mykey2":"ajsae@3if34$/dsd",
    }
    var result string
    var err    error
    for k,v := range pairs {
        suite.doTestSetXattr(k, v)
        result, err = GetXattr(suite.testfile, k)
        suite.Nil(err)
        suite.Equal(result, v)
    }
    result, err = GetXattr(suite.testfile, "myNotExistsTheXKey")
    suite.NotNil(err)
    suite.Equal(result, "")
}

func (suite *UtilsTestSuite) TestDeleteXatrr(){
    //suite.TestSetXattr()
    suite.doTestSetXattr("mykey", "myvalue23")
    err := DeleteXattr(suite.testfile, "mykey")
    suite.Nil(err)
    suite.False(IsXattrExists(suite.testfile, "mykey"))
    err = DeleteXattr(suite.testfile, "myNotExistsTheXKey")
    suite.NotNil(err)
}

func (suite *UtilsTestSuite) TestFileIsExists(){
    result, err := FileIsExists(suite.testfile)
    suite.True(result)
    suite.Nil(err)
    result, err = FileIsExists("#.NoSuchFileExists##!!")
    suite.False(result)
    suite.Nil(err)
}
