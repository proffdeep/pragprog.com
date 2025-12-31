package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"pragprog.com/rggo/interacting/todo"
)

var todoFileName = ".todo.json"

func main() {
	add := flag.Bool("add", false, "Task to include in the list.")
	list := flag.Bool("list", false, "List all stuff.")
	complete := flag.Int("complete", 0, "Is completed?")
	delete := flag.Int("delete", 0, "Delete a task at index.")
	verbose := flag.Bool("verbose",false,"verbosity")
	limit := flag.Int("limit",0,"show only top 10 elements.")
	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {

	case *list:
		if *verbose{
			fmt.Println(time.Now())
		}
		if *limit != 0{
			l.StringN(*limit)
		}else{
			fmt.Print(l)
		}

	case *complete > 0:
		l.Complete(*complete)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _,task := range t{
		l.Add(task) // *(l).Add(item) both are same go handles it internally automatically.
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *delete > 0:
		l.Delete(*delete)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid parameter...")
		os.Exit(1)
	}

}

func getTask(r io.Reader, args ...string) ([]string, error) {
	if len(args) > 0 {
		return []string{strings.Join(args, " ")}, nil
	}

	var tasks []string

	s := bufio.NewScanner(r)
	for s.Scan(){
		text := strings.TrimSpace(s.Text())
		if text != ""{
			tasks = append(tasks,text)
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, fmt.Errorf("Task cannot be blank")
	}
	return tasks, nil
}
