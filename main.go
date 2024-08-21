package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("OI! CUNT! %s This is a PROGRAMMING LANG\n", user.Username)
	fmt.Printf("Type away, yeah\n")
	repl.Start(os.Stdin, os.Stdout)
}
