package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/boxy-pug/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, this is the MOnkey programming language\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
