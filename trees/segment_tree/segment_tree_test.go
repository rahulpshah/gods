package segment_tree

import "testing"
import "github.com/stretchr/testify/assert"

func TestSegTree(t *testing.T) {
	segTree := &SegTree{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	segTree.Build(arr)
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			assert.Equal(t, sum, segTree.Query(segTree.root, i, j))
		}
	}
}
