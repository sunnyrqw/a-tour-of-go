package tour

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestRotReader(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
