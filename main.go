package main

import (
  "fmt"
  "github.com/mstanciu552/todo_go/todo"
  "os"
)

func main() {
  args := os.Args[1:]
  todo.Init()

  if len(args) == 0 {
    todo.ShowAll()
    return
  }

  switch args[1] {

  case "show":
    todo.ShowAll()

  case "add":
    todo.Add()

  case "delete":
    if len(args) == 1 {
      todo.Delete(args[2])
    } else {
      fmt.Println("Id needed")
    }

  case "done":
    if len(args) == 1 {
      todo.Done(args[2])
    } else {
      fmt.Println("Id needed")
    }

  default:
    todo.ShowAll()

  }
}
