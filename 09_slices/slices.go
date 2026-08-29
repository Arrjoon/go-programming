package main

import "fmt"


//slice is dynamic array in python like list 
func main() {
	var abc []int
	fmt.Println(abc==nil)
	fmt.Println(len(abc))
	var nepal = make([]int,2)
	//capacity is maximum numbers of element can fit

	fmt.Println(cap(nepal))
	fmt.Println(nepal)
	// var nums []int
	// fmt.Println(len(nums))
	// fmt.Println(nums==nil)
	//capacity -> maximum number of element can fit 
	var nums = make([]int,2,5)

	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	// nums = append(nums, 5)
	// nums = append(nums, 5)
	// nums := []int{} this is empty slices ;;;;;;;;//////////////////// 

	fmt.Println(nums)
	fmt.Println(cap(nums))
}
