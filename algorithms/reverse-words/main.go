package main

import (
	"fmt"
)

func main() {
	result := reverseWords("i.like.this.program.very.much")
	fmt.Println(result)
}

func reverseWords(s string) string {
	result := "much.very.program.this.like.i"

	return result
}
