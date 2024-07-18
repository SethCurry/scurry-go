package fp_test

import (
	"testing"

	"github.com/SethCurry/scurry-go/fp"
	"github.com/SethCurry/scurry-go/tu"
)

func TestMap(t *testing.T) {
	doubled := fp.Map(func(i int) int {
		return i * 2
	},
		[]int{1, 2, 3, 4, 5})

	tu.SliceEquality[int](t, []int{2, 4, 6, 8, 10}, doubled)
}
