package quicksort

func qsort(data []int) {
	lo := 0
	hi := len(data) - 1

	qs(data, lo, hi)
}

func qs(data []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivotIdx := partition(data, lo, hi)
	qs(data, lo, pivotIdx-1)
	qs(data, pivotIdx+1, hi)
}

func partition(data []int, lo, hi int) int {
	mid := (lo + hi) / 2
	// pivot's from middle and at the end
	tmp := data[mid]
	data[mid] = data[hi]
	data[hi] = tmp

	pivot := data[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if data[i] <= pivot {
			idx++
			tmp := data[i]
			data[i] = data[idx]
			data[idx] = tmp
		}
	}
	idx++

	tmp = data[hi]
	data[hi] = data[idx]
	data[idx] = tmp

	return idx
}
