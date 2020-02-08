package main

import (
	"draíocht/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Gia duit %s! Is é seo an teanga chláreagraithe draíochta!\n",
		user.Username)
	fmt.Printf("Thig leat orduithe a chlóscríobh\n")
	repl.Start(os.Stdin, os.Stdout)
}
