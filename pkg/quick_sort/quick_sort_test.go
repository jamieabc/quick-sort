package quick_sort_test

import (
	"github.com/jamieabc/quick-sort/pkg/quick_sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSorter_Sort(t *testing.T) {
	qs := quick_sort.NewQuickSort([]int{1, 6, 7, 4, 5, 10, 9, 2, 3, 8})
	qs.Sort()

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, qs.Result(), "wrong sort")
}
