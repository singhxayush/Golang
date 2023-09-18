package main

import "fmt"

// array slices and maps

func main() {
	fmt.Println("")
	
// array
	
	// var arr [4]int
	// arr := [4]int{0}
	// arr := [...]int{2, 1, 3, 9}
	
	// var arr [4]uint
	// arr[2] = -1

	// fmt.Println(arr)
	// fmt.Println(arr); fmt.Println(&arr[0]); fmt.Println(&arr[1]); fmt.Println(&arr[2]); fmt.Println(&arr[3]);
	// fmt.Println(arr[1:3])


// slice : https://www.w3schools.com/go/go_slices.php#:~:text=Slices%20are%20similar%20to%20arrays,shrink%20as%20you%20see%20fit.

	// intSlice := []int32{1, 2, 3, 4}
	// intSlice = append(intSlice, 2)
	// fmt.Println(intSlice)

	// var intSlice1 []int32 = []int32{1, 7}
	// fmt.Println(intSlice1)
	// intSlice = append(intSlice, intSlice1...)
	// fmt.Println(intSlice)


	// here, 3, 4 is length and capacity of slice, providing capacity improves performace by avoiding having to reallocate the array
	// var intSlice2 []int32 = make(int32[], 3, 8)


	// slice_name := []datatype{values}

		// sliceName := []string{"Go", "Slices", "Are", "Powerful"}
		// fmt.Println(sliceName)


		// arr1 := [6]int{10, 11, 12, 13, 14,15}
		// myslice := arr1[2:4]
	
		// fmt.Printf("myslice = %v\n", myslice)
		// fmt.Printf("length = %d\n", len(myslice))
		// fmt.Printf("capacity = %d\n", cap(myslice))


// Maps

	var myMap = map[string]int{"Lion":1, "Deer":33, "Fox":12, "Elephant":2}
	fmt.Println(myMap)
	
	// delete(myMap, "Lion")

	fmt.Println(myMap["Lion"])
	fmt.Println(myMap["wosdn"])


	// var mapVal, isPresent = myMap["Lion"]
	// if isPresent {
	// 	fmt.Println("Map Val :", mapVal)
	// } else {
	// 	fmt.Println("Not present")
	// }

	for itr := range myMap {
		// fmt.Println(itr)
		fmt.Printf("%v : %v\n", itr, myMap[itr])
	}

	fmt.Println("")

	for key, val := range myMap {
		fmt.Println(key, " : ",val)
	}

	
}