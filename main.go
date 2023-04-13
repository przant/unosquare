package main

import (
	"fmt"

	"github.com/przant/unosquare/caesar"
)

func main() {
	key := -1
	input := "unosquare is *the* best place to work!"

	if cipher, err := caesar.CaesarCipher(input, key); err != nil {
		fmt.Printf("%s\n\n", err)
	} else {
		for _, char := range cipher {
			fmt.Printf("%s", string(char))
		}
	}
}
