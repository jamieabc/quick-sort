package quick_sort

type Sorter interface {
	Sort()
	Result() []int
}

type sorter struct {
	data []int
}

func (s sorter) Result() []int {
	return s.data
}

func (s *sorter) Sort() {
	length := len(s.data)

	if length <= 1 {
		return
	}

	quickSort(s.data, 0, length-1)
}

// left & right are index
func quickSort(data []int, left, right int) {
	if left < right {
		pivot := partition(data, left, right)
		quickSort(data, left, pivot-1)
		quickSort(data, pivot+1, right)
	}
}

// left & right are index
func partition(data []int, left int, right int) int {
	pivot := data[right]

	var i, j int
	for i, j = left-1, left; j < right; j++ {
		if data[j] < pivot {
			i++
			swap(data, i, j)
		}
	}

	i++
	swap(data, i, right)
	return i
}

func swap(data []int, src, dst int) {
	data[src], data[dst] = data[dst], data[src]
}

func NewQuickSort(data []int) Sorter {
	return &sorter{
		data: data,
	}
}

// problems
// 1. when creating new sorted result, forget about index adjusting, result[i+1] & data[i+2]
