package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

type T interface{ comparable }

func main() {

	values := [][]any{
		{1, true},
		{1, false},
		{1, true},
	}

	deepsort.DeepSort(&values, []any{0})

	for _, e := range values {
		fmt.Println(e)
	}
}
