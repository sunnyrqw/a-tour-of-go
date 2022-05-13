package tour

import (
	"testing"

	"golang.org/x/tour/reader"
)

func TestReader(t *testing.T) {
	reader.Validate(MyReader{})
}
