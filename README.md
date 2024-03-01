# deepsort: Sort slice of slices in Go 

## About
The `deepsort` package provides the `DeepSort` function, which allows sorting a slice of slices based on the values in multiple index positions within the nested slices.

## Installation
```bash
go install github.com/gbatagian/deepsort@latest
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

func main() {
	// Define the slice of slices to be sorted
	data := [][]any{
		{2, "a"},
		{1, "a"},
		{3, "a"},
		{2, "b"},
		{1, "b"},
		{3, "b"},
	}

	// Sort the data based on index positions 0 and 1 (ascending order)
	deepsort.DeepSort(&data, []int{0, 1})

	// Output the sorted data
	for _, row := range data {
		fmt.Println(row)
	}
}
```

Output:
```bash
[1 a]
[1 b]
[2 a]
[2 b]
[3 a]
[3 b]
```

By default, the sorting order is ascending. To specify descending order for an index position, use the negative equivalent of the index position, e.g.:
```go
// Sort based on index position 0 in ascending order and index position 1 in descending order
data = deepsort.DeepSort(&data, []int{0, -1})
```

Based on the above example, this would output:
```bash
[1 b]
[1 a]
[2 b]
[2 a]
[3 b]
[3 a]
```

When sorting in descending order on the zero index position, use `math.Copysign` to force the negative sign. In those cases specify the sort positions as `[]float64` instead of `[]int`, e.g.:
```go
// Sort based on index position 0 in descending order and index position 1 in ascending order
data = deepsort.DeepSort(data, []float64{math.Copysign(0, -1), 1})
```

which would produce the following output based on the initial example:
```bash
[3 a]
[3 b]
[2 a]
[2 b]
[1 a]
[1 b]
```

## Supported Data Types
The slice of slices that is passed as input to the `DeepSort` function can be of any comparable data type, i.e. `[][]int`, `[][]float32`, `[][]string`, `[][]any` etc. As long as the elements at the same index position within the nested slices are of the same type, deepsort can handle them for sorting. `DeepSort` is also designed to be able to handle boolean values when comparing rows. When sorting based on booleans, `false` is considered to be less than `true`. So, when sorting in ascending order, slices with `false` will come before those with `true`, e.g.:

```go
package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

func main() {
	// Define the slice of slices to be sorted
	data := [][]any{
		{3 + 3i, true},
		{3 + 3i, false},
		{1 + 1i, true},
		{1 + 1i, false},
		{2 + 2i, true},
		{2 + 2i, false},
	}

	// Sort the data based on index positions 1 and 0 (ascending order)
	deepsort.DeepSort(&data, []int{1, 0})

	// Output the sorted data
	for _, row := range data {
		fmt.Println(row)
	}
}
```

Which will output:
```bash
[(1+1i) false]
[(2+2i) false]
[(3+3i) false]
[(1+1i) true]
[(2+2i) true]
[(3+3i) true]
```

## Error Handling

If values at the same index position are of different types, the function will panic, e.g.:
```go
package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

func main() {
	// Define the slice of slices to be sorted
	data := [][]any{
		{2, true},
		{2, "false"},
		{1, true},
		{1, false},
	}

	// Sort the data based on index positions 0 and 1 (ascending order)
	deepsort.DeepSort(&data, []int{0, 1})

	// Output the sorted data
	for _, row := range data {
		fmt.Println(row)
	}
}
```
Output:
```bash
panic: Values at the same index position must be of the same type. 
Row: 1 Position: 1  Value: false Type: string
Row: 0 Position: 1  Value: true Type: bool

...
```

However, if an index position with values of different types exists in the nested slices but is not used by `DeepSort`, the algorithm will not panic. For instance, in the above example, if the `DeepSort` call is modified as follows:
```go
// Sort the data based on index positions 0 in ascending order
data = deepsort.DeepSort(&data, []int{0})
```

The output would be:
```bash
[1 true]
[1 false]
[2 true]
[2 false] # This is the row with "false"
```

In this case, the sorting is based only on the first index position, so the rows with mixed types in the second index position do not cause a panic.