package array

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expect := []int{3, 7}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %v want %v", got, expect)
	}
}
