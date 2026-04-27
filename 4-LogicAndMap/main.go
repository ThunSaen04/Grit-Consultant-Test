package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	hashtable := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if _, ok := hashtable[complement]; ok {
			return []int{hashtable[complement], i}
		}
		hashtable[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := TwoSum(nums, target)
	fmt.Println(result)
}
