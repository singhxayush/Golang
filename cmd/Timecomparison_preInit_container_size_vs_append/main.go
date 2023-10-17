package main

import "fmt"
import "time"

func main() {

	fmt.Println("");

	var n int = 100000000
	var test_slice1 = []int{}
	var test_slice2 = make([]int, n)

	fmt.Printf("Time without Pre allocation \t: %v\n", timeLoop(test_slice1, n))
	fmt.Printf("Time with Pre allocation \t: %v\n\n", timeLoop(test_slice2, n))
}

func timeLoop(slice []int, n int) time.Duration {

	var t_start = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t_start)
}

/* 

	Result :
	Time without Pre allocation     : 1.909990545s
	Time with Pre allocation        : 91ns

*/