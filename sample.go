package main

import (
	"fmt"
	"math"

	"github.com/YadaYuki/omochi/pkg/common/slices"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	cnt := 3
	size := int(math.Ceil(float64(len(a)) / float64(cnt)))
	fmt.Println(slices.SplitSlice(a, size))
}
