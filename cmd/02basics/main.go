package main

import "fmt"

// array slices and maps

func main() {

	// array
	// var arr [4]int
	// arr := [4]int{0}
	// arr := [...]int{2, 1, 3, 9}
	
	// var arr [4]uint
	// arr[2] = -1

	// fmt.Println(arr)
	// fmt.Println(arr); fmt.Println(&arr[0]); fmt.Println(&arr[1]); fmt.Println(&arr[2]); fmt.Println(&arr[3]);
	// fmt.Println(arr[1:3])

	// slice
	// intSlice := []int32{1, 2, 3, 4}
	// intSlice = append(intSlice, 2)
	// fmt.Println(intSlice)

	// var intSlice1 []int32 = []int32{1, 7}
	// fmt.Println(intSlice1)
	// intSlice = append(intSlice, intSlice1...)
	// fmt.Println(intSlice)

	// here, 3, 4 is length and capacity of slice, providing capacity improves performace by avoiding having to reallocate the array
	var intSlice2 []int32 = make(int32[], 3, 4)
}