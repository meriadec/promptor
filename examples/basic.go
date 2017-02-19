package main

import (
	"fmt"

	"github.com/meriadec/promptor"
)

func main() {
	items := []string{
		"First choice",
		"Second choice,\non two lines",
		"Third choice",
		"Fourth choice",
	}
	res := promptor.Select(items)
	fmt.Println("result is", res)
}
