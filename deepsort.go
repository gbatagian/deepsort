package deepsort

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"regexp"
	"sort"
	"strconv"
)

type SortPositions interface {
	int | float64 | string
}

// sortConstructor holds the information required for sorting a slice of slices by multiple positions.
type sortConstructor[T comparable] struct {
	data      *[][]T
	positions []any
	reverse   []bool
}

// getValuesFromPosition retrieves values for comparison based on the sort position
// within the nested slices managed by sortConstructor. Positions can be:
//   - Integers (int): referencing an index within the inner slices.
//   - Two-element slices ([2]any):
//     1. First element (int): index within the inner slices.
//     2. Second element (string): field name to extract from a struct at that index.
func (sc *sortConstructor[T]) getValuesFromPosition(pos any, i, j int) (reflect.Value, reflect.Value) {
	switch typedPos := pos.(type) {
	case int:
		valueI := (*sc.data)[i][typedPos]
		valueJ := (*sc.data)[j][typedPos]

		return reflect.ValueOf(valueI), reflect.ValueOf(valueJ)
	case [2]any:
		pos, _ := typedPos[0].(int)
		field, _ := typedPos[1].(string)

		valueI := (*sc.data)[i][pos]
		valueJ := (*sc.data)[j][pos]

		fieldValueI := reflect.ValueOf(valueI).FieldByName(field).Interface().(T)
		fieldValueJ := reflect.ValueOf(valueJ).FieldByName(field).Interface().(T)

		return reflect.ValueOf(fieldValueI), reflect.ValueOf(fieldValueJ)
	default:
		panic(fmt.Errorf("unsupported type %T(%v) for sort position", typedPos, typedPos))
	}
}

// sort sorts a slice of slices based on the values at multiple positions within the inner slices.
//
// Example:
// If sorting by index positions 0 and 1, an input slice of slices like this:
//
//	[][]any{
//		{2, "d"},
//		{2, "c"},
//		{2, "b"},
//		{2, "a"},
//		{1, "d"},
//		{1, "c"},
//		{1, "b"},
//		{1, "a"},
//	}
//
// Would become:
//
//	[][]any{
//		{1, "a"},
//		{1, "b"},
//		{1, "c"},
//		{1, "d"},
//		{2, "a"},
//		{2, "b"},
//		{2, "c"},
//		{2, "d"},
//	}
//
// Note:
// For the sort operation to succeed, the data at the same index position in the nested slices
// must be of the same type. Otherwise, the function will panic.
func (sc *sortConstructor[T]) sort(i, j int) bool {

	XOR := func(a bool, b bool) bool {
		return (a || b) && !(a && b)
	}

	for posIdx, pos := range sc.positions {
		valueI, valueJ := sc.getValuesFromPosition(pos, i, j)

		if valueI.Type() != valueJ.Type() {
			panic(
				fmt.Errorf(
					"sorting error at position %v. Value `%v` and `%v` cannot be compared. Values at the same sort position must be of the same type",
					pos, valueI, valueJ,
				),
			)
		}

		if valueI.Interface() == valueJ.Interface() {
			continue
		}

		switch valueI.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return XOR(valueI.Int() < valueJ.Int(), sc.reverse[posIdx])
		case reflect.Float32, reflect.Float64:
			return XOR(valueI.Float() < valueJ.Float(), sc.reverse[posIdx])
		case reflect.Complex64, reflect.Complex128:
			return XOR(cmplx.Abs(valueI.Complex()) < cmplx.Abs(valueJ.Complex()), sc.reverse[posIdx])
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return XOR(valueI.Uint() < valueJ.Uint(), sc.reverse[posIdx])
		case reflect.String:
			return XOR(valueI.String() < valueJ.String(), sc.reverse[posIdx])
		case reflect.Bool:
			return XOR(!valueI.Bool(), sc.reverse[posIdx])
		}
	}

	return false
}

// DeepSort sorts a slice of slices based on multiple specified positions or fields,
// supporting index-based sorting, field-based sorting within structs, and ascending/descending order.
//
// Supported position types:
//   - int, float64 (converted to int):
//     -- Used as direct index positions for sorting.
//     -- The sign (positive/negative) determines ascending/descending order.
//   - string:
//     -- Facilitates field-based sorting for values retrieved from structs.
//     -- Uses the format "int:fieldName" (e.g., "1:Name" or "-2:Age").
//     -- The int component specifies the index of the struct within the inner slice.
//     -- The string component is the struct field that will be used for sorting.
//     -- The sign on the int component controls ascending/descending order.
func DeepSort[SortPositions, T comparable](sliceOfSlices *[][]T, positions []SortPositions) {

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
			matchRegex, _ := regexp.Compile(`^(\-?[0-9]+):([a-zA-Z]+)$`)
			matchGroups := matchRegex.FindStringSubmatch(typedPos)
			if len(matchGroups) == 0 {
				panic(fmt.Errorf("invalid field specifier format: \"%v\". Use int:fieldName (e.g. \"0:Name\", \"-1:Age\" etc.)", typedPos))
			}

			pos, _ := strconv.Atoi(matchGroups[1])
			field := matchGroups[2]
			if pos < 0 {
				pos = -pos
				sortInReverse[idx] = true
			}
			sortPositions[idx] = [2]any{pos, field}
		default:
			panic(fmt.Errorf("unsupported type %T(%v) provided for sort position. Supported types: int | float64 | string", position, position))
		}
	}

	sc := sortConstructor[T]{data: sliceOfSlices, positions: sortPositions, reverse: sortInReverse}

	sort.Slice(
		*sc.data,
		sc.sort,
	)
}
