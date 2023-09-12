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
	fmt.Println("implementing countsort")
	s := makeRandomSlice(10, 10)
	printSlice(s, 5)
	checkSorted(s)
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

// func countingSort(slice []Customer) {
// 	maxValue := max(slice)
// }
