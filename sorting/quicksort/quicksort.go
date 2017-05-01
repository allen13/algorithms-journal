package main

import "fmt"

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type BySize []int

func (s BySize) Len() int {
	return len(s)
}
func (s BySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s BySize) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	array := []int{8, 6, 7, 5, 3, 0, 9, 8, 9}
	fmt.Println(array)
	Sort(BySize(array))
	fmt.Println(array)
}

//Sort a piece of data
func Sort(data Interface) {
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data Interface, start int, end int) {
	if start < end {
		//Partition array: [lessThan][pivot][greaterThan]
		pivotIndex := partition(data, start, end)

		//quicksort([lessThan])
		quickSort(data, start, pivotIndex-1)

		//quicksort([greaterThan])
		quickSort(data, pivotIndex+1, end)
	}
}

func partition(data Interface, start int, end int) int {
	pivoteIndex := end
	lessThanIterator := start

	//Iterate over the full array range minus the pivot
	// [unsortedValues][pivot]
	for i := start; i <= end-1; i++ {
		if data.Less(i, pivoteIndex) {

			//Move lessThan value into lessThan chunk
			//[lessThan][unsortedValues][pivot]
			data.Swap(lessThanIterator, i)
			lessThanIterator++
		}
	}

	//Move the pivot value in place: [lessThan][greaterThan][pivot] -> [lessThan][pivot][greaterThan]
	data.Swap(lessThanIterator, end)

	return lessThanIterator
}
