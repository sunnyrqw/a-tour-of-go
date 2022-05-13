package tour

import (
	"testing"

	"golang.org/x/tour/pic"
)

func TestImage(t *testing.T) {
	m := Image{100, 100}
	pic.ShowImage(m)

}
