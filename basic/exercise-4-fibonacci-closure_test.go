package tour

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
