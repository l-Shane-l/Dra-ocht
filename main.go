package main

import (
	"fmt"
	"draíocht/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Dia duit %s! agus fáilte go teanga na gclár draíochta!\n",
		user.Username)
	fmt.Printf("tosú ag códú\n")
	repl.Start(os.Stdin, os.Stdout)
}
