package main

import (
	"fmt"
	"os"
)

func main() {
	env, err := ReadDir(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	RunCmd(os.Args[2:], env)
	// fmt.Println()
	// Place your code here.
}
