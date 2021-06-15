package main

import (
  "fmt"
  "go/format"
)

func main() {
  fmt.Println("Formatting Go...")
  src := []byte("func main(){fmt.Println(\"hello\") }")
  fmt.Println(formatGo(src))
}

func formatGo(src []byte) (string, error) {
  bytes, err := format.Source(src)
  if err != nil {
    return "", err
  }
  return string(bytes), nil
}
