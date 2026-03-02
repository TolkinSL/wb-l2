package main

import (
	"fmt"

	"github.com/TolkinSL/wb-l2/l2_9/unpacker"
)

func main() {
	s := `a3bc2`
	// s := `3`
	// s := `a@3bc\2\3`
	// s := `a@3bc\\23`
	// s := `a@3bc\23\`
	// s := `a@3bc\25`
	// s := `a@3\bc\25`
	// s := ""
	u, err := unpacker.Unpack(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s: %s\n", s, u)
}