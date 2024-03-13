# deepsort: Sort slice of slices in Go 

## Contents
1. [About](#about)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Supported Data Types](#supported-data-types)
5. [Error Handling](#error-handling)

## About
The `deepsort` package provides the `DeepSort` function, which allows sorting a slice of slices based on the values in multiple positions within the inner slices. `DeepSort` can handle sorting for any **comparable** type withing the inner slices (such as **integers**, **floats**, **complex numbers**, **strings** and more). It also supports sorting the inner slices based on **struct** types by specifying the struct field to be used for sorting.

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

When sorting the inner slices based on struct fields, you need to provide the combination of the struct's index position within the inner slices and the struct field to be used for sorting in the format `index:FieldName (int:string)`. To specify descending order within this format, utilize the negative equivalent of the index, e.g. `-2:SampleField`. If only struct fields are provided, specify the sort positions as `[]string`. If combinations of index positions with struct fields are provided, use `[]any`, e.g.:
```go
package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

func main() {
	// Define a sample struct
	type Person struct {
		Name string
		Age  int
	}

	// Define the slice of slices to be sorted
	values := [][]any{
		{Person{"Alice", 30}, 1},
		{Person{"Emma", 22}, 2},
		{Person{"Charlie", 18}, 3},
		{Person{"Alice", 42}, 1},
		{Person{"Emma", 37}, 2},
		{Person{"Charlie", 28}, 3},
	}

	// Sort the values based on index position 1 ascending, struct (in index position 0) field "Name" ascending and field "Age" descending
	deepsort.DeepSort(&values, []any{1, "0:Name", "-0:Age"})

	// Output the sorted data
	for _, e := range values {
		fmt.Println(e)
	}
}
```

which would output:
```bash
[{Alice 42} 1]
[{Alice 30} 1]
[{Emma 37} 2]
[{Emma 22} 2]
[{Charlie 28} 3]
[{Charlie 18} 3]
```

## Supported Data Types
The slice of slices that is passed as input to the `DeepSort` function can be of any comparable data type, i.e. `[][]int`, `[][]float32`, `[][]string`, `[][]customStruct`, `[][]any` etc. As long as the elements at the same index position within the nested slices are of the same type, `DeepSort` can handle them for sorting. `DeepSort` is also designed to be able to handle boolean values when comparing the inner slices. When sorting based on booleans, `false` is considered to be less than `true`. So, when sorting in ascending order, slices with `false` will come before those with `true`, e.g.:

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

### Mixed types on the same index position
If values at the same index position are of different types, `DeepSort` will panic, e.g.:
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
panic: sorting error at position 1. Value false (string) and true (bool) cannot be compared. Values at the same sort position must be of the same type
...
```

However, if an index position with values of different types exists in the inner slices but is not used by `DeepSort`, the algorithm will not panic. For instance, in the above example, if the `DeepSort` call is modified as follows:
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

### Invalid string position format
`DeepSort` will panic if a string sort position is not provided in the format `int:string`, e.g. 
```go
package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

func main() {
	// Define a sample struct
	type sample struct {
		Field string
	}

	// Define the slice of slices to be sorted
	values := [][]any{
		{sample{"a"}},
		{sample{"b"}},
		{sample{"c"}},
	}

	// Sort
	deepsort.DeepSort(&values, []any{"--0:Field"})

	// Output the sorted data
	for _, e := range values {
		fmt.Println(e)
	}
}
```

Output:
```bash
panic: invalid field specifier format: "--0:Field". Use int:string (e.g. "0:Name", "-1:Age" etc.)
...
```

### Invalid sort position type
Furthermore, `DeepSort` will panic if the provided sort position is not of type **int**, **float64**, or **string**. For example, if the call in the above example is changed to:
```go
// Sort
deepsort.DeepSort(&values, []any{[]any{0, "Field"}})
```

the function will output:
```bash
panic: unsupported type []interface {}([0 Field]) provided for sort position. Supported types: int | float64 | string
```

### Incomparable sort data types
If the provided data for sorting are incomparable, `DeepSort` wil also panic, e.g.:
```go
package main

import (
	"fmt"

	"github.com/gbatagian/deepsort"
)

func main() {
	// Define the slice of slices to be sorted
	values := [][]any{
		{map[string]int{"a": 1}},
		{map[string]int{"b": 2}},
		{map[string]int{"c": 3}},
	}

	// Sort
	deepsort.DeepSort(&values, []any{0})

	// Output the sorted data
	for _, e := range values {
		fmt.Println(e)
	}
}
```

output:
```bash
panic: runtime error: comparing uncomparable type map[string]int
```