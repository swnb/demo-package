package util

import (
	"testing"
)

func Abc() {

}

type Abcd struct {
}

func TestCompare(t *testing.T) {
	var arr []int
	Compare(1, "a", "", nil, make([]int, 0), arr, Abc, Abcd{}, make([]chan int, 0))
}
