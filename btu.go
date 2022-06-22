package btu

import (
  "io"
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

// Got this from https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file
// Copy the src file to dst. Any existing file will be overwritten and the file
// attributes will not be copied.  Note that like al the other functions in this
// module, an error is fatal to the calling program.
func CopyFile(src, dst string) {
    in, err := os.Open(src)
    CheckError(err)
    defer in.Close()

    out, err := os.Create(dst)
    CheckError(err)
    defer out.Close()

    _, err = io.Copy(out, in)
    CheckError(err)
    out.Close()
}
