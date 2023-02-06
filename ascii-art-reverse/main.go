package main

import (
	"fmt"

	"ascii-art-reverse/funcReverse"
)

func main() {
	charColections := funcReverse.CompileAscii()
	allCharInput := funcReverse.ParseInput()
	// fmt.Println((charColections['{']))
	// fmt.Println((allCharInput[0]))
	result := ""
	for _, eachInput := range allCharInput {
		for i, eachColection := range charColections {
			if eachColection == eachInput {
				result += string(i)
			}
		}
	}
	fmt.Println(result)
}
