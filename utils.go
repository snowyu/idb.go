package iDB

import (
    "os"
    "syscall"
    "xattr"
)

  func ListXattr(aFile string) ([]string, error) {
      keys, err := xattr.Listxattr(aFile)
      var result []string
      if err == nil {
          for _, key := range keys {
              if len(key)> len(XattrPrefix) && key[0:len(XattrPrefix)] == XattrPrefix {
                  result = append(result, key[len(XattrPrefix):])
              }
          }
      }
      return result, err
  }

  func GetXattr(aFile, aKey string) (string, error) {
      value, err := xattr.Getxattr(aFile, XattrPrefix+aKey)
      return string(value), err
  }

  func IsXattrExists(aFile, aKey string) bool {
      return xattr.IsXattrExists(aFile, XattrPrefix+aKey)
  }

  func SetXattr(aFile, aKey, aValue string) error {
      return xattr.Setxattr(aFile, XattrPrefix+aKey, []byte(aValue))
  }

  func DeleteXattr(aFile, aKey string) error {
      return xattr.Removexattr(aFile, XattrPrefix+aKey)
  }

  //check File or Dir is exists
  func FileIsExists(aFile string) (bool, error) {
      _, err := os.Stat(aFile)
      result := err == nil
      if !result {
        e, ok := err.(*os.PathError)
        // Check if error is "no such file or directory"
        if ok && e.Err == syscall.ENOENT {
            err = nil
        }
      }
      return result, err
  }

