package tour

import (
	"fmt"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Printf("========sqrt(%v):%v,math sqrt(%v):%v=========\n\n", i, Sqrt(float64(i)), i, math.Sqrt(float64(i)))
	}
}

func TestSqrtWith1(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Printf("========sqrt(%v):%v,math sqrt(%v):%v=========\n\n", i, SqrtInitWith1(float64(i)), i, math.Sqrt(float64(i)))
	}
}
func TestSqrtWithX(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Printf("========sqrt(%v):%v,math sqrt(%v):%v=========\n\n", i, SqrtInitWithX(float64(i)), i, math.Sqrt(float64(i)))
	}
}

func TestSqrtWithHalfX(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Printf("========sqrt(%v):%v,math sqrt(%v):%v=========\n\n", i, SqrtInitWithHalfX(float64(i)), i, math.Sqrt(float64(i)))
	}
}
