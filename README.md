# deepsort: Sort slice of slices in Go 
# About
The `deepsort` package provides the `DeepSort` function which allows to sort a slice of slices based on the values in multiple index positions in the nested slices. For example, a slice of slices containing the following data:

```go
[][]interface{}{
    {2, "a"},
    {1, "a"},
    {3, "a"},
    {2, "b"},
    {1, "b"},
    {3, "b"},
}
```

if sorted based on the values firstly in the index position `0` and then in the index position `1`, will have the following format after the sort operation:

```go
[][]interface{}{
    {1, "a"},
    {1, "b"},
    {2, "a"},
    {2, "b"},
    {3, "a"},
    {3, "b"},
}
```

The `DeepSort` function can sort in `ascending` and `descening` order, and the order condition can apply in each index position seperately, e.g. in the above example, if the slice of slices is sorted in `ascending` order in the index position `0` and in `descening` order in index position `1`, will eventually have the following format:

```go
[][]interface{}{
    {1, "b"},
    {1, "a"},
    {2, "b"},
    {2, "a"},
    {3, "b"},
    {3, "a"},
}
```

# How to use

Firt install the package
```bash
go get github.com/gbatagian/deepsort
```

Then simply import the package and use the `DeepSort` function, e.g. in a sample `main.go` file
```go
package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

func main() {

	t := [][]interface{}{
		{2, "a"},
		{1, "a"},
		{3, "a"},
		{2, "b"},
		{1, "b"},
		{3, "b"},
	}

	t = deepsort.DeepSort(t, []int{0, 1})

	for _, s := range t {
		fmt.Println(s)
	}
}
```
which will output
```bash
>> go run main.go
[1 a]
[1 b]
[2 a]
[2 b]
[3 a]
[3 b]
```
By default the sort order is `ascending`. To specify a `descending` order in an index position, just pass the equivalent index position with the negative sign infront, e.g.
```go
package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

func main() {

	t := [][]interface{}{
		{2, "a"},
		{1, "a"},
		{3, "a"},
		{2, "b"},
		{1, "b"},
		{3, "b"},
	}

	t = deepsort.DeepSort(t, []int{0, -1})

	for _, s := range t {
		fmt.Println(s)
	}
}
```
which will output
```bash
>> go run main.go
[1 b]
[1 a]
[2 b]
[2 a]
[3 b]
[3 a]
```
To specify `descending` order in the zero index, the `math.Copysign` function needs to be used in order to force the negative sign to the zero value. This means that in that occasion the keys positions indexes will have to be defined as `[]float64`, e.g.
```go
package main

import (
	"fmt"
	"math"

	"github.com/gbatagian/deepsort"
)

func main() {

	t := [][]interface{}{
		{2, "a"},
		{1, "a"},
		{3, "a"},
		{2, "b"},
		{1, "b"},
		{3, "b"},
	}

	t = deepsort.DeepSort(t, []float64{math.Copysign(0, -1), 1})

	for _, s := range t {
		fmt.Println(s)
	}
}
```
which will output
```bash
>> go run main.go
[3 a]
[3 b]
[2 a]
[2 b]
[1 a]
[1 b]
```