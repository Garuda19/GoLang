package sortalgo

type mergeSortStruct struct {
	tempArr []int
}

//NewMergeSortStruct returns sorter for merge sort algo
func NewMergeSortStruct() Sorter {
	return &mergeSortStruct{}
}

func (s *mergeSortStruct) Sort(unsorted []int) {
	s.tempArr = make([]int, len(unsorted))
	s.mergesort(unsorted, 0, len(unsorted)-1)
}

func (s *mergeSortStruct) mergesort(unsorted []int, low, high int) {
	if high <= low {
		return
	}
	mid := (low + high) / 2
	s.mergesort(unsorted, low, mid)
	s.mergesort(unsorted, mid+1, high)
	if unsorted[mid] > unsorted[mid+1] {
		s.merge(unsorted, low, mid, high)
	}
}

func (s *mergeSortStruct) merge(unsorted []int, low, mid, high int) {
	//use i and j as iterators for low to med and high to med
	var i, j int

	for index := low; index <= high; index++ {
		s.tempArr[index] = unsorted[index]
	}

	i = low
	j = mid + 1

	for index := low; index <= high; index++ {
		if i > mid {
			unsorted[index] = s.tempArr[j]
			j++
		} else if j > high {
			unsorted[index] = s.tempArr[i]
			i++
		} else if s.tempArr[i] >= s.tempArr[j] {
			unsorted[index] = s.tempArr[j]
			j++
		} else {
			unsorted[index] = s.tempArr[i]
			i++
		}
	}
}
