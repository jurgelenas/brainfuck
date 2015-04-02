package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "brainfuck/interpreter"
)

func main() {
  i := interpreter.New()

  if len(os.Args[1]) > 0 {
    src, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
      fmt.Println("Error reading file")
      fmt.Println(err)
    } else {
      i.Load(string(src))
      output := i.Run()

      fmt.Print(output)
    }
  } else {
    fmt.Println("Usage: ./brainfuck filepath/code")
  }
}
