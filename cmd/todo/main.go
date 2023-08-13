package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/piyushyadav0191/cmd-todo"
)

const (
	todoFile = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark to do as completed")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	
	flag.Parse() 

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil{
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample Todo")
		err := todos.Store(todoFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		}

	case *complete> 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		}

	case *del> 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		}		

		case *list:
			todos.Print()

	default:
		fmt.Fprintln(os.Stdout, "Invalid command")
		os.Exit(0)
	}
}