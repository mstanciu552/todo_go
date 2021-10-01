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

	switch args[0] {

	case "clear":
		todo.ClearDB()

	case "show":
		todo.ShowAll()

	case "add":
		if args[1] != "" {
			todo.Add(args[1]) // title
		} else {
			fmt.Println("Id needed")
		}

  case "due":
    todo.Due(args[1], args[2])

	case "delete":
		if args[1] != "" {
			todo.Delete(args[1]) // id
		} else {
			fmt.Println("Id needed")
		}

	case "done":
		if args[1] != "" {
			todo.Done(args[1]) // id
		} else {
			fmt.Println("Id needed")
		}

	default:
		todo.ShowAll()

	}
}
