package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println("Hello World!")
	// var intNum int16 = 32767
	// var intNum uint16 = 32767
	// var intNum int16 = 32767 + 1 // -> given an error !!!
	// intNum++
	// fmt.Println(intNum)

	// var strVar string = "AbC"
	// fmt.Println(strVar)

	// x := 1
	// fmt.Println(x)

	// var a, b = 1, 2
	// println(a, b)

	// const pi float32 = 3.14
	// println(pi)

	// var printVar string = "Hello Fucking World!"
	// print_this(printVar)

	// println(divide(3, 2))

	var res, rem, err = intDivide(7, 0)
	// if err!=nil{
	// 	fmt.Println(err.Error())
	// }
	// var res, rem int = intDivide(7, 0) // panic: runtime error: integer divide by zero
	// fmt.Println(res, rem)
	// fmt.Printf("res = %v | rem = %v", res, rem)

	
	switch{
	case err != nil:
		fmt.Println(err.Error())
	case rem == 0:
		fmt.Printf("res = %v", res)
	default:
		fmt.Printf("res = %v | rem = %v", res, rem)
	}

	// switch rem{
	// case 0:
	// 	fmt.Printf("division was exact")
	// case 1, 2:
	// 	fmt.Printf("The div was close")
	// default:
	// 	fmt.Print("the div was not close")
	// }
}

// functions
func print_this(printVar string) {
	fmt.Println(printVar)
}

// with return type
func divide(num int, den int) int {
	return num / den
}

// with multiple return types
func intDivide(num int, den int) (int, int, error) {
	// error has default val as nil
	var err error
	if den == 0 {
		err = errors.New("cant divide by 0")
		return 0, 0, err
	}
	return num / den, num % den, err
}
