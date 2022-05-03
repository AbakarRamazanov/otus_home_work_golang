package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func reverse(s string) string {
	return stringutil.Reverse(s)
}

func main() {
	fmt.Println(reverse("Hello, OTUS!"))
	// Place your code here.
}
