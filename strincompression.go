package main

import (
	"fmt"
)

func stringCompression(str string) (output string) {
	strmap := make(map[rune]int)
	runes := []rune(str)
	for _, ru := range runes {
		strmap[ru]++
	}

	for key, value := range strmap {
		output += fmt.Sprintf("%s%d", string(key), value)

	}
	return output
}
func main() {
	fmt.Println(stringCompression("aaabbcccccdddeeeefffeee"))
}
