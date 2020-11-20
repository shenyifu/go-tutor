package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 5)
	expect := "aaaaa"
	if repeat != expect {
		t.Errorf("expect '%q' , but got '%q'", expect, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("ab", 5)
	}
}
