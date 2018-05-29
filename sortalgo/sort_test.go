package sortalgo

import (
	"fmt"
	"math/rand"
	"testing"
)

func getRandomArray() []int {
	randomLength := 10000
	randomArray := make([]int, randomLength)
	for i := 0; i < randomLength; i++ {
		randomArray[i] = rand.Intn(10000)
	}
	return randomArray
}

func sorted(sortedarr []int) bool {
	isSorted := true
	for i := range sortedarr {
		if i > 0 && sortedarr[i] < sortedarr[i-1] {
			isSorted = false
		}
	}
	return isSorted
}

func testSort(s Sorter, t *testing.T) {
	array := getRandomArray()
	s.Sort(array)
	fmt.Println(array)
	if !sorted(array) {
		t.Error("Array not sorted")
	}
}
func TestQuickSort(t *testing.T) {
	s := NewQuickSortStruct()
	testSort(s, t)
}

func TestMergeSort(t *testing.T) {
	s := NewMergeSortStruct()
	testSort(s, t)
}
