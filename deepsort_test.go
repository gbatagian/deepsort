package deepsort

import (
	"math"
	"testing"
)

func TestDeepSort(t *testing.T) {

	rawValues := [][]any{
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

	rawValues = DeepSort(&rawValues, []int{0, 2, 1, 3})

	expectedSortedRawValues := [][]any{
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

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}

}

func TestDeepSortReverse(t *testing.T) {

	rawValues := [][]any{
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

	rawValues = DeepSort(&rawValues, []int{0, -2, 1, -3})

	expectedSortedRawValues := [][]any{
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

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}

}

func TestDeepSortReverseInZeroIndexStart(t *testing.T) {

	rawValues := [][]any{
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

	rawValues = DeepSort(&rawValues, []float64{math.Copysign(0, -1), 3})

	expectedSortedRawValues := [][]any{
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

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}

}

func TestDeepSortReverseInZeroIndexMiddle(t *testing.T) {

	rawValues := [][]any{
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

	rawValues = DeepSort(&rawValues, []float64{2, math.Copysign(0, -1), 3})

	expectedSortedRawValues := [][]any{
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

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}

}

func TestDeepSortComplexNumbers(t *testing.T) {
	rawValues := [][]any{
		{1, false, 1 + 1i},
		{1, false, 2 + 2i},
		{1, false, 3 + 3i},
		{2, true, 1 + 1i},
		{2, true, 2 + 2i},
		{2, true, 3 + 3i},
	}

	rawValues = DeepSort(&rawValues, []int{2, 0, 1})

	expectedSortedRawValues := [][]any{
		{1, false, (1 + 1i)},
		{2, true, (1 + 1i)},
		{1, false, (2 + 2i)},
		{2, true, (2 + 2i)},
		{1, false, (3 + 3i)},
		{2, true, (3 + 3i)},
	}

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortComplexNumbersReverse(t *testing.T) {
	rawValues := [][]any{
		{1, false, 1 + 1i},
		{1, false, 2 + 2i},
		{1, false, 3 + 3i},
		{2, true, 1 + 1i},
		{2, true, 2 + 2i},
		{2, true, 3 + 3i},
	}

	rawValues = DeepSort(&rawValues, []float64{-2, math.Copysign(0, -1), -1})

	expectedSortedRawValues := [][]any{
		{2, true, (3 + 3i)},
		{1, false, (3 + 3i)},
		{2, true, (2 + 2i)},
		{1, false, (2 + 2i)},
		{2, true, (1 + 1i)},
		{1, false, (1 + 1i)},
	}

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}
}

func TestDeepSortUnsignedIntegers(t *testing.T) {
	rawValues := [][]any{
		{uint(1), false, 1 + 1i},
		{uint(1), false, 2 + 2i},
		{uint(1), false, 3 + 3i},
		{uint(2), true, 1 + 1i},
		{uint(2), true, 2 + 2i},
		{uint(2), true, 3 + 3i},
	}

	rawValues = DeepSort(&rawValues, []float64{math.Copysign(0, -1), 2})

	expectedSortedRawValues := [][]any{
		{uint(2), true, (1 + 1i)},
		{uint(2), true, (2 + 2i)},
		{uint(2), true, (3 + 3i)},
		{uint(1), false, (1 + 1i)},
		{uint(1), false, (2 + 2i)},
		{uint(1), false, (3 + 3i)},
	}

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
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

	rawValues := [][]any{
		{1, false},
		{1, "false"},
	}

	DeepSort(&rawValues, []int{1})

}
