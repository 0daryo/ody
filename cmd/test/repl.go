package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/0daryo/ody/repl"
)

func Exec() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the ody programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
