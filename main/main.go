package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

type T interface{ comparable }

func main() {
	type sample struct {
		Field any
	}

	values := [][]any{
		{1, sample{"a"}},
		{1, sample{"b"}},
	}

	deepsort.DeepSort(&values, []any{"-1:Field"})

	fmt.Println(values)
}
