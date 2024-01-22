// main.go

package main

import (
	"fmt"
	"orangutan/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, this is the Orangutan interactive repl.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
