package main

import (
	"fmt"
	"math/rand"
)

type Customer struct {
	id           string
	numPurchases int
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}

func makeRandomSlice(numItems, max int) []Customer {
	s := make([]Customer, numItems)
	for i := 0; i < numItems; i++ {
		s[i].id = fmt.Sprintf("C%d", i)
		s[i].numPurchases = rand.Intn(max)
	}
	return s
}

func printSlice(slice []Customer, numItems int) {
	if numItems > len(slice) {
		numItems = len(slice)
	}
	s := slice[:numItems]
	fmt.Printf("%v \n", s)
}

func checkSorted(slice []Customer) {
	sorted := true
	msg := "The slice is sorted"
	for i := 0; i < (len(slice) - 1); i++ {
		if slice[i].numPurchases > slice[i+1].numPurchases {
			sorted = false
			break
		}
	}
	if !sorted {
		msg = "The slice is NOT sorted!"
	}
	fmt.Println(msg)
}

// func maxValue(slice []Customer) int {
// 	max := slice[0].numPurchases
// 	for _, v := range slice {
// 		if v.numPurchases > max {
// 			max = v.numPurchases
// 		}
// 	}
// 	return max
// }

func countingSort(slice []Customer, max int) []Customer {
	// Create a big enough counts slice
	countsSlice := make([]int, max+1)
	// Fill up the counts slice
	for _, v := range slice {
		countsSlice[v.numPurchases]++
	}
	// rescale counts Slice
	for i := 1; i < len(countsSlice); i++ {
		countsSlice[i] += countsSlice[i-1]
	}
	// create a new Customer slice
	workingSlice := make([]Customer, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		k := slice[i].numPurchases
		workingSlice[countsSlice[k]-1] = slice[i]
		countsSlice[k]--
	}
	// copy(slice, workingSlice)
	return workingSlice
}
