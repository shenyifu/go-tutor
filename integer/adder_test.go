package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4
	if sum != expect {
		t.Errorf("expect '%d' but got '%d'", expect, sum)
	}
}

func ExampleAdder() {
	sum := Add(1, 3)
	fmt.Println(sum)
	// Output: 4
}
