package deepsort

import (
	"math"
	"testing"
)

func TestDeepSort(t *testing.T) {
	values := [][]any{
		{1, .2, false, "b"},
		{1, .2, false, "a"},
		{1, .1, false, "h"},
		{2, .1, false, "d"},
		{2, .1, true, "f"},
		{2, .1, false, "e"},
		{1, .2, false, "g"},
		{1, .1, false, "c"},
		{1, .2, true, ""},
	}

	DeepSort(&values, []int{0, 2, 1, 3})

	sortedValues := [][]any{
		{1, 0.1, false, "c"},
		{1, 0.1, false, "h"},
		{1, 0.2, false, "a"},
		{1, 0.2, false, "b"},
		{1, 0.2, false, "g"},
		{1, 0.2, true, ""},
		{2, 0.1, false, "d"},
		{2, 0.1, false, "e"},
		{2, 0.1, true, "f"},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortReverse(t *testing.T) {
	values := [][]any{
		{1, .2, false, "b"},
		{1, .2, false, "a"},
		{1, .1, false, "h"},
		{2, .1, false, "d"},
		{2, .1, true, "f"},
		{2, .1, false, "e"},
		{1, .2, false, "g"},
		{1, .1, false, "c"},
		{1, .2, true, ""},
	}

	DeepSort(&values, []int{0, -2, 1, -3})

	sortedValues := [][]any{
		{1, 0.2, true, ""},
		{1, 0.1, false, "h"},
		{1, 0.1, false, "c"},
		{1, 0.2, false, "g"},
		{1, 0.2, false, "b"},
		{1, 0.2, false, "a"},
		{2, 0.1, true, "f"},
		{2, 0.1, false, "e"},
		{2, 0.1, false, "d"},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortReverseInZeroIndexStart(t *testing.T) {
	values := [][]any{
		{1, .2, false, "b"},
		{1, .2, false, "a"},
		{1, .1, false, "h"},
		{2, .1, false, "d"},
		{2, .1, true, "f"},
		{2, .1, false, "e"},
		{1, .2, false, "g"},
		{1, .1, false, "c"},
		{1, .2, true, ""},
	}

	DeepSort(&values, []float64{math.Copysign(0, -1), 3})

	sortedValues := [][]any{
		{2, 0.1, false, "d"},
		{2, 0.1, false, "e"},
		{2, 0.1, true, "f"},
		{1, 0.2, true, ""},
		{1, 0.2, false, "a"},
		{1, 0.2, false, "b"},
		{1, 0.1, false, "c"},
		{1, 0.2, false, "g"},
		{1, 0.1, false, "h"},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortReverseInZeroIndexMiddle(t *testing.T) {

	values := [][]any{
		{1, .2, false, "b"},
		{1, .2, false, "a"},
		{1, .1, false, "h"},
		{2, .1, false, "d"},
		{2, .1, true, "f"},
		{2, .1, false, "e"},
		{1, .2, false, "g"},
		{1, .1, false, "c"},
		{1, .2, true, ""},
	}

	DeepSort(&values, []float64{2, math.Copysign(0, -1), 3})

	sortedValues := [][]any{
		{2, 0.1, false, "d"},
		{2, 0.1, false, "e"},
		{1, 0.2, false, "a"},
		{1, 0.2, false, "b"},
		{1, 0.1, false, "c"},
		{1, 0.2, false, "g"},
		{1, 0.1, false, "h"},
		{2, 0.1, true, "f"},
		{1, 0.2, true, ""},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}

}

func TestDeepSortComplexNumbers(t *testing.T) {
	values := [][]any{
		{1, false, 1 + 1i},
		{1, false, 2 + 2i},
		{1, false, 3 + 3i},
		{2, true, 1 + 1i},
		{2, true, 2 + 2i},
		{2, true, 3 + 3i},
	}

	DeepSort(&values, []int{2, 0, 1})

	sortedValues := [][]any{
		{1, false, (1 + 1i)},
		{2, true, (1 + 1i)},
		{1, false, (2 + 2i)},
		{2, true, (2 + 2i)},
		{1, false, (3 + 3i)},
		{2, true, (3 + 3i)},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortComplexNumbersReverse(t *testing.T) {
	values := [][]any{
		{1, false, 1 + 1i},
		{1, false, 2 + 2i},
		{1, false, 3 + 3i},
		{2, true, 1 + 1i},
		{2, true, 2 + 2i},
		{2, true, 3 + 3i},
	}

	DeepSort(&values, []float64{-2, math.Copysign(0, -1), -1})

	sortedValues := [][]any{
		{2, true, (3 + 3i)},
		{1, false, (3 + 3i)},
		{2, true, (2 + 2i)},
		{1, false, (2 + 2i)},
		{2, true, (1 + 1i)},
		{1, false, (1 + 1i)},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortUnsignedIntegers(t *testing.T) {
	values := [][]any{
		{uint(1), false, 1 + 1i},
		{uint(1), false, 2 + 2i},
		{uint(1), false, 3 + 3i},
		{uint(2), true, 1 + 1i},
		{uint(2), true, 2 + 2i},
		{uint(2), true, 3 + 3i},
	}

	DeepSort(&values, []float64{math.Copysign(0, -1), 2})

	sortedValues := [][]any{
		{uint(2), true, (1 + 1i)},
		{uint(2), true, (2 + 2i)},
		{uint(2), true, (3 + 3i)},
		{uint(1), false, (1 + 1i)},
		{uint(1), false, (2 + 2i)},
		{uint(1), false, (3 + 3i)},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortValueOnTheSameIndexPositionNotSameType(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Did not panic.")
		}
	}()

	values := [][]any{
		{1, false},
		{1, "false"},
	}

	DeepSort(&values, []int{1})
}

func TestDeepSortInt(t *testing.T) {
	values := [][]int64{
		{10, 20, 30, 40},
		{100, 200, 300, 400},
		{1000, 2000, 3000, 4000},
	}

	DeepSort(&values, []float64{math.Copysign(0, -1)})

	sortedValues := [][]int64{
		{1000, 2000, 3000, 4000},
		{100, 200, 300, 400},
		{10, 20, 30, 40},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortString(t *testing.T) {
	values := [][]string{
		{"apple", "d"},
		{"banana", "c"},
		{"cherry", "b"},
		{"orange", "a"},
		{"elderberry", "e"},
	}

	DeepSort(&values, []int{1})

	sortedValues := [][]string{
		{"orange", "a"},
		{"cherry", "b"},
		{"banana", "c"},
		{"apple", "d"},
		{"elderberry", "e"},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortComplex(t *testing.T) {
	values := [][]complex64{
		{1 + 1i},
		{2 + 2i},
		{3 + 3i},
	}

	DeepSort(&values, []float64{math.Copysign(0, -1)})

	sortedValues := [][]complex64{
		{3 + 3i},
		{2 + 2i},
		{1 + 1i},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortFloat(t *testing.T) {
	values := [][]float64{
		{3.14159},
		{1.61803},
		{2.71828},
	}

	DeepSort(&values, []int{0})

	sortedValues := [][]float64{
		{1.61803},
		{2.71828},
		{3.14159},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortEmpty(t *testing.T) {
	values := [][]any{}

	DeepSort(&values, []int{0})

	sortedValues := [][]any{}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortStruct(t *testing.T) {
	type sample struct {
		Field string
	}

	values := [][]any{
		{sample{"b"}},
		{sample{"a"}},
		{sample{"b"}},
		{sample{"a"}},
	}

	DeepSort(&values, []any{"0:Field"})

	sortedValues := [][]any{
		{sample{"a"}},
		{sample{"a"}},
		{sample{"b"}},
		{sample{"b"}},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortStructReverse(t *testing.T) {
	type sample struct {
		Field string
	}

	values := [][]any{
		{2, sample{"a"}},
		{2, sample{"b"}},
		{1, sample{"a"}},
		{1, sample{"b"}},
	}

	DeepSort(&values, []any{0, "-1:Field"})

	sortedValues := [][]any{
		{1, sample{"b"}},
		{1, sample{"a"}},
		{2, sample{"b"}},
		{2, sample{"a"}},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortStructReverseZeroPosition(t *testing.T) {
	type sample struct {
		Field string
	}

	values := [][]any{
		{sample{"a"}, 1},
		{sample{"b"}, 1},
		{sample{"a"}, 2},
		{sample{"b"}, 2},
	}

	DeepSort(&values, []any{"-0:Field", -1})

	sortedValues := [][]any{
		{sample{"b"}, 2},
		{sample{"b"}, 1},
		{sample{"a"}, 2},
		{sample{"a"}, 1},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortSliceOfSlicesOfStructsInput(t *testing.T) {
	type sample struct {
		Field string
	}

	values := [][]sample{
		{sample{"a"}},
		{sample{"b"}},
		{sample{"a"}},
		{sample{"b"}},
	}

	DeepSort(&values, []any{"0:Field"})

	sortedValues := [][]sample{
		{sample{"a"}},
		{sample{"a"}},
		{sample{"b"}},
		{sample{"b"}},
	}

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestNoSwapsOnEqualRows(t *testing.T) {
	values := [][]any{
		{1, true},
		{1, false},
		{1, true},
	}

	DeepSort(&values, []any{0})

	sortedValues := values

	for sIDx, s := range sortedValues {
		for vIdx, v := range s {
			if !(v == values[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortStructPositionFalseFormat1(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Did not panic.")
		}
	}()

	type sample struct {
		Field string
	}

	values := [][]any{
		{sample{"a"}, 1},
		{sample{"b"}, 1},
		{sample{"a"}, 2},
		{sample{"b"}, 2},
	}

	DeepSort(&values, []any{"--0:Field"})

}

func TestDeepSortStructPositionFalseFormat2(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Did not panic.")
		}
	}()

	type sample struct {
		Field string
	}

	values := [][]any{
		{sample{"a"}, 1},
		{sample{"b"}, 1},
		{sample{"a"}, 2},
		{sample{"b"}, 2},
	}

	DeepSort(&values, []any{"-0::Field"})

}

func TestDeepSortStructPositionFalseFormat3(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Did not panic.")
		}
	}()

	type sample struct {
		Field string
	}

	values := [][]any{
		{sample{"a"}, 1},
		{sample{"b"}, 1},
		{sample{"a"}, 2},
		{sample{"b"}, 2},
	}

	DeepSort(&values, []any{":Field"})

}

func TestDeepSortStructPositionFalseFormat4(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Did not panic.")
		}
	}()

	type sample struct {
		Field string
	}

	values := [][]any{
		{sample{"a"}, 1},
		{sample{"b"}, 1},
		{sample{"a"}, 2},
		{sample{"b"}, 2},
	}

	DeepSort(&values, []any{"0a:Field"})

}

func TestDeepSortStructPositionFalseFormat5(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Did not panic.")
		}
	}()

	type sample struct {
		Field string
	}

	values := [][]any{
		{sample{"a"}, 1},
		{sample{"b"}, 1},
		{sample{"a"}, 2},
		{sample{"b"}, 2},
	}

	DeepSort(&values, []any{" 0:Field"})

}

func TestDeepSortStructPositionFalseFormat6(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Did not panic.")
		}
	}()

	type sample struct {
		Field string
	}

	values := [][]any{
		{sample{"a"}, 1},
		{sample{"b"}, 1},
		{sample{"a"}, 2},
		{sample{"b"}, 2},
	}

	DeepSort(&values, []any{"0:Field "})

}
