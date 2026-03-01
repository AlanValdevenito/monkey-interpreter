package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/AlanValdevenito/monkey-interpreter/monkey"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user information:", err)
		return
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n\n", user.Username)
	fmt.Printf("Feel free to type in commands\n\n")

	m := monkey.New(os.Stdin, os.Stdout)

	if len(os.Args) > 1 {
		scriptFile := os.Args[1]
		err := m.RunScriptFile(scriptFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running script: %v\n", err)
		}
		return
	}

	m.StartREPL()
}