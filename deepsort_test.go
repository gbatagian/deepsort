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

	rawValues = DeepSort(rawValues, []int{0, 2, 1, 3})

	expectedSortedRawValues := [][]any{
		{1, .2, true, ""},
		{1, .1, false, "c"},
		{1, .1, false, "h"},
		{1, .2, false, "a"},
		{1, .2, false, "b"},
		{1, .2, false, "g"},
		{2, .1, true, "f"},
		{2, .1, false, "d"},
		{2, .1, false, "e"},
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

	rawValues = DeepSort(rawValues, []int{0, -2, 1, -3})

	expectedSortedRawValues := [][]any{
		{1, .1, false, "h"},
		{1, .1, false, "c"},
		{1, .2, false, "g"},
		{1, .2, false, "b"},
		{1, .2, false, "a"},
		{1, .2, true, ""},
		{2, .1, false, "e"},
		{2, .1, false, "d"},
		{2, .1, true, "f"},
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

	rawValues = DeepSort(rawValues, []float64{math.Copysign(0, -1), 3})

	expectedSortedRawValues := [][]any{
		{2, .1, false, "d"},
		{2, .1, false, "e"},
		{2, .1, true, "f"},
		{1, .2, true, ""},
		{1, .2, false, "a"},
		{1, .2, false, "b"},
		{1, .1, false, "c"},
		{1, .2, false, "g"},
		{1, .1, false, "h"},
	}

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}

}

func TestDeepSortReverseInZeroIndexMidle(t *testing.T) {

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

	rawValues = DeepSort(rawValues, []float64{2, math.Copysign(0, -1), 3})

	expectedSortedRawValues := [][]any{
		{2, .1, true, "f"},
		{1, .2, true, ""},
		{2, .1, false, "d"},
		{2, .1, false, "e"},
		{1, .2, false, "a"},
		{1, .2, false, "b"},
		{1, .1, false, "c"},
		{1, .2, false, "g"},
		{1, .1, false, "h"},
	}

	for sIDx, s := range expectedSortedRawValues {
		for vIdx, v := range s {
			if !(v == rawValues[sIDx][vIdx]) {
				t.Error("Unexpected data resulted after sort operation")
			}
		}
	}

}
