package main

import (
	"fmt"

	"github.com/gophertrain/trainctl/cmd"
)

var Version string

func main() {
	fmt.Printf("trainctl %s\n", Version)
	cmd.Execute()
}
