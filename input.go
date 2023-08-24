package btu

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func GetNumberOrQuit(prompt string) int {
  for {
    s := GetResponse(prompt)
    if len(s) < 1 {
      continue
    }
    if s[0] == 'q' {
      os.Exit(0)
    }
    n, err := strconv.Atoi(s)
    if err != nil {
      fmt.Printf("Can't convert '%s' to a number, try again.\n", s)
      continue
    }
    return n
  }
}

func GetResponse(prompt string) string {
  fmt.Printf("%s: ", prompt)
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return strings.TrimSpace(scanner.Text())
}
