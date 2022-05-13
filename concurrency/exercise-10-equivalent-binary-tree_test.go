package tour

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	time.Sleep(time.Second)
}

func TestSameTree(t *testing.T) {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
