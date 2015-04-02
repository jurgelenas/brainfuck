package interpreter

import (
  "os"
)

type Interpreter struct {
  memory []uint8
  location int
  program string
  output string
}

const MEM_SIZE = 300

func New() Interpreter {
  i := Interpreter{}
  return i
}

func (i *Interpreter) Load(program string) {
  i.memory = make([]uint8, MEM_SIZE);
  i.location = 0
  i.program = program
  i.output = ""
}

func (i *Interpreter) Run() string {
  for instruction := 0; instruction < len(i.program); instruction++ {
    switch i.program[instruction] {
      case '>':
        i.location += 1
        if len(i.memory) <= i.location {
          i.memory = append(i.memory, 0)
        }

      case '<':
        if i.location > 0 {
          i.location -= 1
        }

      case '+':
        i.memory[i.location]++

      case '-':
        i.memory[i.location]--

      case '.':
        i.output += string(i.memory[i.location])

      case ',':
        b := make([]byte, 1)
        os.Stdin.Read(b)
        i.memory[i.location] = b[0]

      case '[':
        if i.memory[i.location] == 0 {
          instruction++;
          loop_depth := 0;
          for loop_depth > 0 || i.program[instruction] != ']' {
            if i.program[instruction] == '[' {
              loop_depth++
            } else if i.program[instruction] == ']' {
              loop_depth--
            }
            instruction++;
          }
        }

      case ']':
        instruction--;
        loop_depth := 0;
        for loop_depth > 0 || i.program[instruction] != '[' {
          if i.program[instruction] == ']' {
            loop_depth++
          } else if i.program[instruction] == '[' {
            loop_depth--
          }
          instruction--;
        }
        instruction--;
    }
  }

  return i.output
}
