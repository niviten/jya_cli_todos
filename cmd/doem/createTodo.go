package main

import (
	"doem/todos"
	"flag"
	"fmt"
	"time"
)

func createTodo(args []string) {
    f := flag.NewFlagSet("create", flag.ExitOnError)

	var name string
	var description string
	var dueDate string
	var priority int

	f.StringVar(&name, "name", "default name", "Todo name")
	f.StringVar(&description, "desc", "default description", "Todo description")
	f.StringVar(&dueDate, "due", "default due date", "Due date")
	f.IntVar(&priority, "priority", 5, "Priority")

    f.Parse(args)

	fmt.Printf("subcommand = %s\n", "create")
	fmt.Printf("name = %s\n", name)
	fmt.Printf("description = %s\n", description)
	fmt.Printf("due date = %s\n", dueDate)
	fmt.Printf("priority = %d\n", priority)

    if len(name) == 0 {
        fmt.Println("Name can't be empty")
        return
    }
    
    if priority < 1 || priority > 5 {
        fmt.Println("Priority should be in between [1, 5]")
        return
    }

    todo := todos.New()
    todo.Name = name
    todo.Description = description
    todo.DueDate = dueDate
    todo.Priority = priority

    todo.CreatedDate = time.Now().Format("2006-01-02")
}
