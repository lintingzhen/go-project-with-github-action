package main

import (
	"fmt"
	"go-project-with-github-action/sub"
)

var version = "0.0"

func main() {
	fmt.Printf("version=%s\n", version)
	sub.Print()
}
