package btu

import (
  "log"
  "os"
)

func FileExists(path string) bool {
  fileInfo, err := os.Stat(path)
  if err != nil {
    return false
  }
  if !fileInfo.Mode().IsRegular() {
    log.Fatal("%s exists, but is not a file\n", path)
  }
  return true
}

func DirExists(dir string) bool {
  fileInfo, err := os.Stat(dir)
  if err != nil {
    return false
  }
  if !fileInfo.IsDir() {
    log.Fatal("%s exists, but is not a directory\n", dir)
  }
  return true
}

func DirMustExist(dir string) {
  if !DirExists(dir) {
    log.Fatal("Directory %s does not exist\n", dir)
  }
}

func CheckError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func CreateFile(path string) *os.File {
  file, err := os.Create(path)
  CheckError(err)
  return file
}
