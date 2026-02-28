package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/AlanValdevenito/monkey-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user information:", err)
		return
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n\n", user.Username)
	fmt.Printf("Feel free to type in commands\n\n")
	repl.Start(os.Stdin, os.Stdout)
}