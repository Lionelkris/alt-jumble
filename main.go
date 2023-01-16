package main

import (
	"fmt"
	"os"

	"github.com/Lionelkris/alt-jumble/pkg/altstrjumble"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("ERROR: Two arguments exactly are mandatory for the application")
		os.Exit(1)
	}
	s := altstrjumble.AltJumble(args[1], args[2])
	fmt.Printf("The jumbled string is %s\n", s)
}
