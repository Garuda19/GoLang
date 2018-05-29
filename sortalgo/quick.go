package sortalgo

type quickSortStruct struct{}

//NewQuickSortStruct returns pointer to atruct for implementing quicksort
func NewQuickSortStruct() Sorter {
	return &quickSortStruct{}
}

func (s *quickSortStruct) quickSort(unsorted []int, low, high int) {
	if high <= low {
		return
	}
	pivot := s.partition(unsorted, low, high)
	s.quickSort(unsorted, low, pivot-1)
	s.quickSort(unsorted, pivot+1, high)
}

func (s *quickSortStruct) Sort(unsorted []int) {
	s.quickSort(unsorted, 0, len(unsorted)-1)
}

func (s *quickSortStruct) partition(unsorted []int, low, high int) int {
	i := low
	j := high
	pivot := low
	for {
		for {
			if i >= high || unsorted[i] > unsorted[pivot] {
				break
			}
			i++
		}
		for {
			if j <= low || unsorted[j] < unsorted[pivot] {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		Swap(unsorted, i, j)
	}
	Swap(unsorted, pivot, j)
	return j
}
