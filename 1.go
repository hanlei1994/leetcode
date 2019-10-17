package main

import "fmt"

func main() {
	fmt.Println("leetcode 1 ")
	//	intls := []int{2, 4, 5, 7, 8, 9, 0}
	intls := []int{3, 2, 4}
	fmt.Println(twoSum(intls, 6))

}

// func twoSum(nums []int, target int) []int {

// 	index := make(map[int]int, len(nums))

// 	for i, b := range nums {

// 		if j, ok := index[target-b]; ok {

// 			return []int{j, i}
// 		}
// 		index[b] = i
// 	}
// 	return nil
// }

func twoSum(nums []int, target int) []int {

	for i, b := range nums {
		for j, a := range nums {
			if i == j {
				break
			}
			if a+b == target {
				fmt.Println(i, j)

				return []int{j, i}

			}

		}
	}
	return nil

}
