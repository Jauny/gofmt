package main

import (
  "go/format"
  "syscall/js"
)

func main() {
  c := make(chan struct{})
  registerCallback()
  <-c
}

func registerCallback() {
  js.Global().Set("formatFromJs", js.FuncOf(formatFromJs))
}

func formatFromJs(this js.Value, inputs []js.Value) interface{} {
  rawGoCode := inputs[0].String()
  formatted, err := formatGo([]byte(rawGoCode))
  if err != nil {
    return err
  }

  return formatted
}

func formatGo(src []byte) (string, error) {
  bytes, err := format.Source(src)
  if err != nil {
    return "", err
  }
  return string(bytes), nil
}
