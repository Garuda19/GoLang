package sortalgo

//Sorter is interface for every sorting algo
type Sorter interface {
	Sort(unsorted []int)
}

//GetMin returns minimum of 2 int
func GetMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//Swap swaps elements at indexes a & b of given array
func Swap(arr []int, a, b int) {
	k := arr[a]
	arr[a] = arr[b]
	arr[b] = k
}
