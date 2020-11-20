package array

import (
	"reflect"
	"testing"
)

func TestSumAllTail(t *testing.T) {

	checkSum := func(t *testing.T, want, got []int) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	}
	t.Run("first test case", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{3, 7})
		expect := []int{2, 7}
		checkSum(t, expect, got)
	})

	t.Run("empty test", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 5, 1})
		want := []int{0, 6}
		checkSum(t, want, got)
	})

}
