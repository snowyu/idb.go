package iDB

import (
    //"os"
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

