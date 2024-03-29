package btu

import (
  "fmt"
  "io"
  "log"
  "os"
  "path/filepath"
  "strconv"
  "strings"
)

func FileExists(path string) bool {
  fileInfo, err := os.Stat(path)
  if err != nil {
    return false
  }
  if !fileInfo.Mode().IsRegular() {
    log.Fatalf("%s exists, but is not a file\n", path)
  }
  return true
}

func DirExists(dir string) bool {
  fileInfo, err := os.Stat(dir)
  if err != nil {
    return false
  }
  if !fileInfo.IsDir() {
    log.Fatalf("%s exists, but is not a directory\n", dir)
  }
  return true
}

// Like the above, but doesn't fatal out if the path is a regular file.
func IsDir(dir string) bool {
  fileInfo, err := os.Stat(dir)
  if err != nil {
    return false
  }
  return fileInfo.IsDir()
}

func DirMustExist(dir string) {
  if !DirExists(dir) {
    log.Fatalf("Directory %s does not exist\n", dir)
  }
}

func CreateDir(dir string) {
  err := os.MkdirAll(dir, 0755)
  CheckError(err)
}

func CreateDirForFile(path string) {
  dir, _ := filepath.Split(path)
  CreateDir(dir)
}

func CheckError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func CheckError2(err error, msg string, a ...any) {
  if err != nil {
    m := fmt.Sprintf(msg, a...)
    Fatal("%s\n%s\n", m, err.Error())
  }
}

func CreateFile(path string) *os.File {
  file, err := os.Create(path)
  CheckError(err)
  return file
}

func OpenFile(path string) *os.File {
  file, err := os.Open(path)
  CheckError(err)
  return file
}

// The B suffix means "byte".
func ReadFileB(path string) []byte {
  b, err := os.ReadFile(path)
  CheckError(err)
  return b
}

// The S suffix means "string".
func ReadFileS(path string) string {
  return string(ReadFileB(path))
}

// The L suffix means "lines".
func ReadFileL(path string) []string {
  s := ReadFileS(path)
  return strings.Split(s, "\n")
}

// Got this from https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file
// Copy the src file to dst. Any existing file will be overwritten and the file
// attributes will not be copied.  Note that like all the other functions in this
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

func Atoi(a string) int {
  j, err := strconv.Atoi(a)
  CheckError(err)
  return j
}

func Atoi2(s string, msg string, a ...any) int {
  j, err := strconv.Atoi(s)
  CheckError2(err, msg, a...)
  return j
}

func ParseInt64(s string, msg string, a ...any) int64 {
  j, err := strconv.ParseInt(s, 10, 64)
  CheckError2(err, msg, a...)
  return j
}

func ParseFloat64(s string, msg string, a ...any) float64 {
  f, err := strconv.ParseFloat(s, 64)
  CheckError2(err, msg, a...)
  return f
}
