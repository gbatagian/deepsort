package deepsort

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"regexp"
	"slices"
	"strconv"
)

type SortPosition interface {
	int | float64 | string
}

// sortConstructor holds the information required for sorting a slice of slices by multiple positions.
type sortConstructor[T comparable] struct {
	data      *[][]T
	positions []any
	reverse   []bool
}

// getValuesFromPosition extracts the values for comparison based on the specified sort position
// within the input slices (a, b []T). These extracted values are utilized by the [sortConstructor.sort]
// method to determine the sort order between the two slices.
//
// Positions (pos any) can be:
//   - Integers (int): Represent indices for directly referencing elements within the input slices (a, b []T).
//   - Two-element slices ([2]any):
//     1. First element (int): Index referencing the position of a struct within the input slices (a, b []T).
//     2. Second element (string): Field name to extract from the struct at the specified index.
func (sc *sortConstructor[T]) getValuesFromPosition(pos any, a, b []T) (reflect.Value, reflect.Value) {
	switch typedPos := pos.(type) {
	case int:
		valueI := a[typedPos]
		valueJ := b[typedPos]

		return reflect.ValueOf(valueI), reflect.ValueOf(valueJ)
	case [2]any:
		pos, _ := typedPos[0].(int)
		field, _ := typedPos[1].(string)

		valueA := a[pos]
		valueB := b[pos]

		return reflect.ValueOf(valueA).FieldByName(field), reflect.ValueOf(valueB).FieldByName(field)
	default:
		panic(fmt.Errorf("unsupported type %T(%v) for sort position", typedPos, typedPos))
	}
}

// compare sorts a slice of slices based on the values at multiple positions within the inner slices.
//
// Example:
// If sorting by index positions 0 and 1, an input slice of slices like this:
//
//	[][]any{
//		{2, "b"},
//		{2, "a"},
//		{1, "b"},
//		{1, "a"},
//	}
//
// Would become:
//
//	[][]any{
//		{1, "a"},
//		{1, "b"},
//		{2, "a"},
//		{2, "b"},
//	}
//
// Note:
// For the compare operation to succeed, the data at the same index position in the (a, b []T) slices
// must be of the same type. Otherwise, the function will panic.
func (sc *sortConstructor[T]) compare(a, b []T) int {

	xorIntMap := func(b1 bool, b2 bool) int {
		// Implements the XOR gate and maps true to 1 and false to -1.
		// This map aligns with the expected return type (int) of the comparison
		// function required by [slices.SortFunc] (to which the compare method is passed).
		//
		// Example:
		// Consider the use of xorIntMap in the statements
		//
		//		return xorIntMap(valueA.Int() > valueB.Int(), sc.reverse[posIdx])
		//
		// The comparison logic `valueA > valueB` is combined with the `reverse` flag using XOR and
		// the final output determined the sort order:
		//
		// |  valueA > valueB  |  reverse  |  xor outcome  | mapped value  |
		// |     (b1 bool)     | (b2 bool) |      ---      |   (return)    |
		// |        T          |    T      |       F       |      -1       |  -> move b to lower index (swap)
		// |        F          |    T      |       T       |       1       |  -> leave a and b as is (no swap)
		// |        T          |    F      |       T       |       1       |  -> leave a and b as is (no swap)
		// |        F          |    F      |       F       |      -1       |  -> move b to lower index (swap)
		if (b1 || b2) && !(b1 && b2) {
			return 1
		}
		return -1
	}

	for posIdx, pos := range sc.positions {
		valueA, valueB := sc.getValuesFromPosition(pos, a, b)

		if valueA.Type() != valueB.Type() {
			panic(
				fmt.Errorf(
					"sorting error at position %v. Value %v (%v) and %v (%v) cannot be compared. Values at the same sort position must be of the same type",
					pos, valueA, valueA.Kind(), valueB, valueB.Kind(),
				),
			)
		}

		if valueA.Interface() == valueB.Interface() {
			// If values are equal in a position, go to the next position in order
			// to determine the the sort order of (a, b []T).
			continue
		}

		switch valueA.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return xorIntMap(valueA.Int() > valueB.Int(), sc.reverse[posIdx])
		case reflect.Float32, reflect.Float64:
			return xorIntMap(valueA.Float() > valueB.Float(), sc.reverse[posIdx])
		case reflect.Complex64, reflect.Complex128:
			return xorIntMap(cmplx.Abs(valueA.Complex()) > cmplx.Abs(valueB.Complex()), sc.reverse[posIdx])
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return xorIntMap(valueA.Uint() > valueB.Uint(), sc.reverse[posIdx])
		case reflect.String:
			return xorIntMap(valueA.String() > valueB.String(), sc.reverse[posIdx])
		case reflect.Bool:
			return xorIntMap(valueA.Bool(), sc.reverse[posIdx])
		}
	}
	return 0
}

// DeepSort sorts a slice of slices based on multiple user-specified positions or fields.
// It supports various sorting criteria inputs:
//
//   - Index-based sorting (int/float, converted to int):
//     -- Used for direct element access within the inner slices for sorting.
//     -- Sign determines ascending/descending order (positive/negative).
//   - Field-based sorting (structs):
//     -- Use strings in the format "index:fieldName" (e.g., "1:Name" or "-2:Age")
//     -- The integer specifies the index of the struct within the inner slice.
//     -- The field name identifies the struct field used for sorting.
//     -- The sign on the integer controls ascending/descending order.
//
// Index-based sorting and Field-based sorting can be used (if needed) in combination, e.g.:
//
//	type sample struct {
//		Field string
//	}
//
//	values := [][]any{
//		{2, sample{"a"}},
//		{2, sample{"b"}},
//		{1, sample{"a"}},
//		{1, sample{"b"}},
//	}
//
//	deepsort.DeepSort(&values, []any{0, "-1:Field"})
//
// which will result `values` to be:
//
//	[1 {b}]
//	[1 {a}]
//	[2 {b}]
//	[2 {a}]
//
// Alternatively `positions []SortPositions` argument can be []int or []float64.
// Refer README.md for more example and use cases.
func DeepSort[SortPosition, T comparable](sliceOfSlices *[][]T, positions []SortPosition) {

	sortPositions := make([]any, len(positions))
	sortInReverse := make([]bool, len(positions))

	for idx, position := range positions {
		switch pos := interface{}(position); pos.(type) {
		case float64:
			typedPos, _ := pos.(float64)
			sortPositions[idx] = int(typedPos)
			if math.Signbit(typedPos) {
				sortPositions[idx] = -int(typedPos)
				sortInReverse[idx] = true
			}
		case int:
			typedPos, _ := pos.(int)
			sortPositions[idx] = typedPos
			if typedPos < 0 {
				sortPositions[idx] = -typedPos
				sortInReverse[idx] = true
			}
		case string:
			typedPos := pos.(string)
			matchRegex, _ := regexp.Compile(`^(\-?[0-9]+):([a-zA-Z0-9_]+)$`)
			matchGroups := matchRegex.FindStringSubmatch(typedPos)
			if len(matchGroups) == 0 {
				panic(fmt.Errorf("invalid field specifier format: \"%v\". Use int:string (e.g. \"0:Name\", \"-1:Age\" etc.)", typedPos))
			}

			pos, _ := strconv.Atoi(matchGroups[1])
			field := matchGroups[2]
			if string(matchGroups[1][0]) == "-" {
				pos = -pos
				sortInReverse[idx] = true
			}
			sortPositions[idx] = [2]any{pos, field}
		default:
			panic(fmt.Errorf("unsupported type %T(%v) provided for sort position. Supported types: int | float64 | string", position, position))
		}
	}

	sc := sortConstructor[T]{data: sliceOfSlices, positions: sortPositions, reverse: sortInReverse}

	slices.SortFunc(
		*sc.data,
		sc.compare,
	)
}
