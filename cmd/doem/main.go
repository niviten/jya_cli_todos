package main

import (
	"fmt"
	"os"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }

    subcommand := os.Args[1]

    switch subcommand {
    case "create":
        createTodo(os.Args[2:])
    default:
        fmt.Printf("Unknown command: %s\n", subcommand)
        printUsage()
        os.Exit(1)
    }

    fmt.Println("___done")
}

func printUsage() {
    fmt.Println("Usage: mona <command> [--name <name>]" +
        "[--desc <description>] [--due <YYYY-MM-DD>] [--priority <1-5>]")
}
