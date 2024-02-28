package deepsort

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"sort"
)

type sortable interface {
	uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 | complex64 | complex128 |
		string | bool | any
}

// sortConstructor holds the information required for sorting a slice of slices by multiple index positions.
type sortConstructor[s sortable] struct {
	data      *[][]s
	positions []int
	reverse   []bool
}

// sortSliceByPositions sorts a slice of slices based on the values at multiple index positions.
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
// of the slice of slices must be of the same type. Otherwise, the function will panic.
func (sc *sortConstructor[sortable]) sortSliceByPositions(i, j int) bool {

	XOR := func(a bool, b bool) bool {
		return (a || b) && !(a && b)
	}

	for posIdx, pos := range sc.positions {
		v1 := reflect.ValueOf((*sc.data)[i][pos])
		v2 := reflect.ValueOf((*sc.data)[j][pos])

		if v1.Type() != v2.Type() {
			panic(
				"Values at the same index position must be of the same type. " +
					fmt.Sprintf("\nRow: %v Position: %v  Value: %v Type: %T", i, pos, v1, (*sc.data)[i][pos]) +
					fmt.Sprintf("\nRow: %v Position: %v  Value: %v Type: %T", j, pos, v2, (*sc.data)[j][pos]),
			)
		}

		if v1.Interface() == v2.Interface() {
			continue
		}

		switch v1.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return XOR(v1.Int() < v2.Int(), sc.reverse[posIdx])
		case reflect.Float32, reflect.Float64:
			return XOR(v1.Float() < v2.Float(), sc.reverse[posIdx])
		case reflect.Complex64, reflect.Complex128:
			return XOR(cmplx.Abs(v1.Complex()) < cmplx.Abs(v2.Complex()), sc.reverse[posIdx])
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return XOR(v1.Uint() < v2.Uint(), sc.reverse[posIdx])
		case reflect.String:
			return XOR(v1.String() < v2.String(), sc.reverse[posIdx])
		case reflect.Bool:
			return XOR(!v1.Bool(), sc.reverse[posIdx])
		}
	}

	return false
}

func DeepSort[idxPosition int | float64, s sortable](sliceOfSlices *[][]s, positions []idxPosition) {

	sortPositions := make([]int, len(positions))
	sortInReverse := make([]bool, len(positions))

	for idx, idxPos := range positions {
		if idxPos == 0 {
			if math.Signbit(float64(idxPos)) {
				sortInReverse[idx] = true
			}
			sortPositions[idx] = 0
		} else if idxPos < 0 {
			sortPositions[idx] = -int(idxPos)
			sortInReverse[idx] = true
		} else {
			sortPositions[idx] = int(idxPos)
		}
	}

	sc := sortConstructor[s]{data: sliceOfSlices, positions: sortPositions, reverse: sortInReverse}

	sort.Slice(
		*sc.data,
		sc.sortSliceByPositions,
	)
}
