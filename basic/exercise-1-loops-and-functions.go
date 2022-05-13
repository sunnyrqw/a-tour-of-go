package tour

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / 2 * z
		fmt.Printf("z:%v\n", z)
	}
	return z
}

func SqrtFunction(x float64, z float64) float64 {
	var result float64
	for i := 0; i < 20; i++ {
		z, result = z-(z*z-x)/2*z, z
		fmt.Println(z)
		if math.Abs(result-z) < 1e-3 {
			break
		}
	}
	return z
}

func SqrtInitWith1(x float64) float64 {
	return SqrtFunction(x, 1)
}

func SqrtInitWithX(x float64) float64 {
	return SqrtFunction(x, x)
}

func SqrtInitWithHalfX(x float64) float64 {
	return SqrtFunction(x, x/2)
}
