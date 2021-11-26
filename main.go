package main

import (
	"fmt"
	"mortar/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the repl of Mortar programming language!\n", user.Username)
	fmt.Printf("Type 'help' to see the list of commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
