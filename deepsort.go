package deepsort

import (
	"fmt"
	"math"
	"sort"
)

// Private struct which holds the sort operation related information
type sortConstructor struct {
	data                        *[][]interface{}
	sortKeysPositions           []int
	reverse                     []bool
	currentSortKeyPositionIndex int
}

// sortSliceWithMultipleIndexPositions sorts a slice of slices by mulpiple index positions, e.g. the slice:

//		[][]interface{}{
//			{2, "d"},
//			{2, "c"},
//			{2, "b"},
//			{2, "a"},
//			{1, "d"},
//			{1, "c"},
//			{1, "b"},
//			{1, "a"},
//		}

// if sorted by index positions 0, 1 to become:

//		[][]interface{}{
//			{1, "a"},
//			{1, "b"},
//			{1, "c"},
//			{1, "d"},
//			{2, "a"},
//			{2, "b"},
//			{2, "c"},
//			{2, "d"},
//		}

// sortSliceWithMultipleIndexPositions executes the sort operation by starting comparing the rows based on the first specified
// index position, and if 2 rows are equal based on the current index, it does a recursive call and applied the sort operation
// in the next specified index position. The recursion is executed until either an inequality result is achived or there are no
// more available index position, which means that the 2 rows compared are equal and no swap should take place.

// !! Note: for the sort operation to be succesfull, the data on the same index position in the nested slices of the slice of slices
// !!       should be of the type, else the function will panic.

func (sc *sortConstructor) sortSliceWithMultipleIndexPositions(i, j int) bool {

	if sc.currentSortKeyPositionIndex > len(sc.sortKeysPositions)-1 {
		// This if condition is to end the recursive calls.
		// To end up in here it means that two rows are equal based on the checks in all the specified key index positions
		sc.currentSortKeyPositionIndex = 0 // renormilise sort key position index (to perform a fresh check in the next rows)
		return true
	}

	idx := sc.sortKeysPositions[sc.currentSortKeyPositionIndex]
	var con bool

	data := *sc.data
	switch v1 := data[i][idx].(type) {
	case int:
		v2, ok := data[j][idx].(int)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 < v2 {
			con = true
		} else if v1 > v2 {
			con = false
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	case int8:
		v2, ok := data[j][idx].(int8)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 < v2 {
			con = true
		} else if v1 > v2 {
			con = false
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	case int16:
		v2, ok := data[j][idx].(int16)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 < v2 {
			con = true
		} else if v1 > v2 {
			con = false
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	case int32:
		v2, ok := data[j][idx].(int32)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 < v2 {
			con = true
		} else if v1 > v2 {
			con = false
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	case int64:
		v2, ok := data[j][idx].(int64)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 < v2 {
			con = true
		} else if v1 > v2 {
			con = false
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	case float32:
		v2, ok := data[j][idx].(float32)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 < v2 {
			con = true
		} else if v1 > v2 {
			con = false
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	case float64:
		v2, ok := data[j][idx].(float64)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 < v2 {
			con = true
		} else if v1 > v2 {
			con = false
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	case string:
		v2, ok := data[j][idx].(string)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 < v2 {
			con = true
		} else if v1 > v2 {
			con = false
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	case bool:
		v2, ok := data[j][idx].(bool)
		if !ok {
			panic(fmt.Sprintf("Can not compare values %v and %v in index position %d in nested slices %d and %d", v1, data[j][idx], idx, i+1, j+1))
		}
		if v1 != v2 {
			con = v1
		} else {
			// rows are equal based on the current key index, so evaluate inequality condition in the next key index
			sc.currentSortKeyPositionIndex += 1
			return sc.sortSliceWithMultipleIndexPositions(i, j)
		}
	default:
		con = false
	}

	if sc.reverse[sc.currentSortKeyPositionIndex] {
		con = !con
	}

	if sc.currentSortKeyPositionIndex > 0 {
		// renormilise sort key position index after each inequality condition is resolved (to perform a fresh check in the next rows)
		sc.currentSortKeyPositionIndex = 0
	}

	return con

}

func DeepSort[kIdx int | float64](s [][]interface{}, k []kIdx) [][]interface{} {

	keysPositions := make([]int, len(k))
	sortInReverseOrderMap := make([]bool, len(k))
	for idx, keyIndex := range k {
		if keyIndex == 0 {
			zeroValue := float64(keyIndex)
			if math.Signbit(zeroValue) {
				sortInReverseOrderMap[idx] = true
			}
			keysPositions[idx] = 0
		} else if keyIndex < 0 {
			keysPositions[idx] = -int(keyIndex)
			sortInReverseOrderMap[idx] = true
		} else {
			keysPositions[idx] = int(keyIndex)
		}
	}

	sc := sortConstructor{data: &s, sortKeysPositions: keysPositions, reverse: sortInReverseOrderMap}

	sort.Slice(
		*sc.data,
		sc.sortSliceWithMultipleIndexPositions,
	)

	return *sc.data

}
