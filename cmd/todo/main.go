package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/piyushyadav0191/cmd-todo"
)

const (
	todoFile = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add a new work")
	complete := flag.Int("complete", 0, "mark work as completed")
	del := flag.Int("del", 0, "delete a work")
	list := flag.Bool("list", false, "list all works")

	
	flag.Parse() 

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil{
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
	task, err := getInput(os.Stdin, flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)   
	}
	todos.Add(task)
err = todos.Store(todoFile)
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

func getInput(r io.Reader, args ...string) (string, error){
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", nil
	}
	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("Empty work is not allowed")
	}

	return text, nil
}