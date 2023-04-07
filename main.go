package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ehsaaniqbal/notc/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("This is not C, go nuts %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
