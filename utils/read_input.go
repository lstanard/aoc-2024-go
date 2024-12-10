package utils

import (
  "fmt"
  "io"
  "os"
  "strings"
)

func ReadInputFile(path string) (bool, []byte) {
  fmt.Println("Reading input file:", path)

  file, err := os.Open(path)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return false, nil
  }
  defer file.Close()

  content, err := io.ReadAll(file)
  if err != nil {
    fmt.Println("Error reading file:", err)
    return false, nil
  }

  return true, content
}

func SplitInputToLines(data []byte) []string {
  stringData := string(data)
  lines := strings.Split(stringData, "\n")
  return lines
}
