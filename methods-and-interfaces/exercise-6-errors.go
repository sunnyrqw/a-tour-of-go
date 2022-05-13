package tour

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//compile error: fmt.Sprintf format %v with arg e causes recursive (xxx.ErrNegativeSqrt).Error method call
	//return fmt.Sprintf("cannot Sqrt negative number:%v", e)
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)

	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / 2 * z
		fmt.Printf("z:%v\n", z)
	}
	return z, nil
}
